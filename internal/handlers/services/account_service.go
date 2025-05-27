package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"server/internal/handlers/dto"
	repository "server/internal/handlers/repositories"
	model "server/internal/models"
	"server/internal/utils/dotenv"
	CustomHash "server/internal/utils/hash"
	jwtutil "server/internal/utils/jwt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"gorm.io/gorm"
)

type AccountService struct {
	accountRepo repository.AccountRepository
}

func NewAccountService(accountRepo repository.AccountRepository) *AccountService {
	return &AccountService{accountRepo: accountRepo}
}

type CustomClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func init() {

}

func (s *AccountService) GetFullAccountWithEmail(email string) ([]*model.Account, error) {
	accounts, err := s.accountRepo.GetByEmail(email)
	if err != nil {
		return accounts, err
	}
	return accounts, nil
}
func (s *AccountService) GetFullAccountWithUser(userId string) ([]*model.Account, error) {
	accounts, err := s.accountRepo.GetByUserId(userId)
	if err != nil {
		return accounts, err
	}
	return accounts, nil
}

func (s *AccountService) Register(account *model.Account) error {
	accounts, err := s.accountRepo.GetByEmail(account.Email)
	if err == nil {
		_, err := s.accountRepo.GetByEmailAndMethod(account.Email, account.LoginMethod)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				hash, err := CustomHash.HashPassword(account.Password)
				if err != nil {
					return errors.New("cannot hash the password")
				}
				var lastLogin *time.Time
				var role string
				if len(accounts) > 0 {
					lastLogin = accounts[0].LastLogin
					role = accounts[0].Role
				} else {
					lastLogin = nil
					role = "student"
				}
				Account := &model.Account{
					AccountId:   CustomHash.HashMD5(time.Now().String()),
					UserId:      account.UserId,
					Email:       account.Email,
					Password:    string(hash),
					Role:        role,
					Status:      "offline",
					LastLogin:   lastLogin,
					LoginMethod: account.LoginMethod,
					CreatedAt:   account.CreatedAt,
					UpdatedAt:   account.UpdatedAt,
				}
				return s.accountRepo.Create(Account)
			} else {
				return errors.New("oops!! Server is error. Sorry")
			}
		} else {
			log.Println("User existing!")
			return errors.New("email already registered")
		}
	} else {
		return errors.New("can't connect to server")
	}
}

func (s *AccountService) Login(input *model.Account) (map[string]string, error) {
	Account, err := s.accountRepo.GetByEmailAndMethod(input.Email, input.LoginMethod)
	if err != nil || Account == nil {
		return nil, errors.New("invalid email or password")
	}

	if !CustomHash.CheckPassword(Account.Password, input.Password) {
		return nil, errors.New("password is incorrect")
	}
	Account.Status = "online"
	if err := s.accountRepo.Update(Account); err != nil {
		return nil, errors.New("server update account encountered an error")
	}

	accessToken, err := jwtutil.CreateToken(Account.UserId, jwtutil.GetAccessToken(), jwtutil.GetAccessExpire())
	if err != nil {
		return nil, err
	}
	refreshToken, err := jwtutil.CreateToken(Account.UserId, jwtutil.GetRefreshToken(), jwtutil.GetRefreshExpire())
	if err != nil {
		return nil, err
	}
	var last_login string
	LastLogin := Account.LastLogin
	if LastLogin == nil {
		last_login = ""
	} else {
		last_login = LastLogin.String()
	}

	return map[string]string{
		"access_token":    accessToken,
		"refresh_token":   refreshToken,
		"user_id":         Account.UserId,
		"role":            Account.Role,
		"last_login":      last_login,
		"access_exprise":  time.Now().Add(jwtutil.GetAccessExpire()).String(),
		"refresh_exprise": time.Now().Add(jwtutil.GetRefreshExpire()).String(),
	}, nil
}

func (s *AccountService) Logout(userId *string) error {
	Accounts, err := s.accountRepo.GetByUserId(*userId)
	if err != nil || Accounts == nil {
		return errors.New("invalid email or password")
	}
	now := time.Now()
	for _, acc := range Accounts {
		acc.LastLogin = &now
		acc.Status = "offline"

		if err := s.accountRepo.Update(acc); err != nil {
			return errors.New("cannot logout")
		}
	}
	return nil
}

func (s *AccountService) LoginWithGoogle(input *model.Account) (map[string]string, error) {
	Account, err := s.accountRepo.GetByEmailAndMethod(input.Email, input.LoginMethod)
	if err != nil || Account.AccountId == "" {
		Account.UserId = input.UserId
		Account.CreatedAt = time.Now()
		Account.UpdatedAt = time.Now()
		Account.LastLogin = input.LastLogin
		Account.AccountId = CustomHash.HashMD5(time.Now().String())
		Account.Email = input.Email
		Account.Role = input.Role
		Account.Status = "online"
		Account.LoginMethod = "google"
		if err := s.accountRepo.Create(Account); err != nil {
			return nil, errors.New("server update account encountered an error")
		}
	} else {
		Account.Status = "online"
		if err := s.accountRepo.Update(Account); err != nil {
			return nil, errors.New("server update account encountered an error")
		}
	}

	accessToken, err := jwtutil.CreateToken(Account.UserId, jwtutil.GetAccessToken(), jwtutil.GetAccessExpire())
	if err != nil {
		return nil, err
	}
	refreshToken, err := jwtutil.CreateToken(Account.UserId, jwtutil.GetRefreshToken(), jwtutil.GetRefreshExpire())
	if err != nil {
		return nil, err
	}
	var last_login string
	LastLogin := Account.LastLogin
	if LastLogin == nil {
		last_login = ""
	} else {
		last_login = LastLogin.String()
	}

	return map[string]string{
		"access_token":    accessToken,
		"refresh_token":   refreshToken,
		"user_id":         Account.UserId,
		"role":            Account.Role,
		"last_login":      last_login,
		"access_exprise":  time.Now().Add(jwtutil.GetAccessExpire()).String(),
		"refresh_exprise": time.Now().Add(jwtutil.GetRefreshExpire()).String(),
	}, nil
}

func (s *AccountService) LoginWithGithub(state *dto.GitHubUserRequest) string {
	var githubOAuthConfig = &oauth2.Config{
		ClientID:     dotenv.GetDotEnv("GITHUB_CLIENT_ID"),
		ClientSecret: dotenv.GetDotEnv("GITHUB_CLIENT_SECRET"),
		RedirectURL:  dotenv.GetDotEnv("APP_URL") + ":" + dotenv.GetDotEnv("APP_PORT") + "/api/auth/github/callback",
		Scopes:       []string{"user", "user:email"},
		Endpoint:     github.Endpoint,
	}
	claims := jwt.MapClaims{
		"redirect_uri": state.RedirectURI,
		"nonce":        state.Nonce,
		"exp":          state.ExpiresAt,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedState, err := token.SignedString([]byte(jwtutil.GetAccessToken()))
	if err != nil {
		return ""
	}

	url := githubOAuthConfig.AuthCodeURL(signedState)
	return url
}

func (s *AccountService) GitHubCallback(code string) (*dto.GitHubUserResponse, error) {
	githubOAuthConfig := &oauth2.Config{
		ClientID:     dotenv.GetDotEnv("GITHUB_CLIENT_ID"),
		ClientSecret: dotenv.GetDotEnv("GITHUB_CLIENT_SECRET"),
		RedirectURL:  dotenv.GetDotEnv("APP_URL") + ":" + dotenv.GetDotEnv("APP_PORT") + "/api/auth/github/callback",
		Scopes:       []string{"user", "user:email"},
		Endpoint:     github.Endpoint,
	}

	token, err := githubOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange code for token: %v", err)
	}

	client := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(token))

	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %v", err)
	}
	defer resp.Body.Close()

	var user dto.GitHubUserResponse
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("failed to decode user info: %v", err)
	}

	respEmail, err := client.Get("https://api.github.com/user/emails")
	if err != nil {
		return nil, fmt.Errorf("failed to get user emails: %v", err)
	}
	defer respEmail.Body.Close()

	var emails []struct {
		Email    string `json:"email"`
		Primary  bool   `json:"primary"`
		Verified bool   `json:"verified"`
	}
	if err := json.NewDecoder(respEmail.Body).Decode(&emails); err != nil {
		return nil, fmt.Errorf("failed to decode email info: %v", err)
	}

	for _, e := range emails {
		if e.Primary && e.Verified {
			user.Email = e.Email
			break
		}
	}
	return &user, nil
}

func (ac *AccountService) GetAccountByEmailAndMethod(email string, method string) (*model.Account, error) {
	return ac.accountRepo.GetByEmailAndMethod(email, method)
}
func (ac *AccountService) GetAccountsByUser(userId string) ([]*model.Account, error) {
	return ac.accountRepo.GetByUserId(userId)
}

func (ac *AccountService) Create(account *model.Account) error {
	return ac.accountRepo.Create(account)
}
func (ac *AccountService) Update(account *model.Account) bool {
	if err := ac.accountRepo.Update(account); err != nil {
		return false
	}
	return true
}
func (ac *AccountService) Delete(accountId string) error {
	return ac.accountRepo.Delete(accountId)
}
