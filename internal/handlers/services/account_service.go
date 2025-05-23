package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"server/internal/handlers/dto"
	repository "server/internal/handlers/repositories"
	model "server/internal/models"
	CustomHash "server/internal/utils/hash"
	jwtutil "server/internal/utils/jwt"
	"time"

	"github.com/golang-jwt/jwt/v4"
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
	log.Println("Register")
	_, err := s.accountRepo.GetByEmailAndMethod(account.Email, account.LoginMethod)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hash, err := CustomHash.HashPassword(account.Password)
			if err != nil {
				return errors.New("cannot hash the password")
			}
			Account := &model.Account{
				AccountId:   CustomHash.HashMD5(time.Now().String()),
				UserId:      account.UserId,
				Email:       account.Email,
				Password:    string(hash),
				Role:        "student",
				Status:      "offline",
				LastLogin:   nil,
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
}

func (s *AccountService) Login(input *model.Account) (map[string]string, error) {
	Account, err := s.accountRepo.GetByEmailAndMethod(input.Email, input.LoginMethod)
	if err != nil || Account == nil {
		return nil, errors.New("invalid email or password")
	}

	if CustomHash.CheckPassword(Account.Password, input.Password) {
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

func (ac *AccountService) Update(account *model.Account) bool {
	if err := ac.accountRepo.Update(account); err != nil {
		return false
	}
	return true
}

func (s *AccountService) LoginWithGoogle(input *model.Account) (map[string]string, error) {
	Account, err := s.accountRepo.GetByEmailAndMethod(input.Email, input.LoginMethod)
	if err != nil || Account.AccountId == "" {
		Account.UserId = input.UserId
		Account.CreatedAt = time.Now()
		Account.UpdatedAt = time.Now()
		Account.LastLogin = nil
		Account.AccountId = CustomHash.HashMD5(time.Now().String())
		Account.Email = input.Email
		Account.Role = "student"
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

func (s *AccountService) LoginWithGitHub(accessToken string) (*dto.GitHubUserResponse, error) {
	if accessToken == "" {
		return nil, errors.New("access token is empty")
	}

	url := "https://api.github.com/user"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("token %s", accessToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to get github user info: %s", string(bodyBytes))
	}

	var user dto.GitHubUserResponse
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return nil, err
	}

	// chưa thêm login vào đây này!
	return &user, nil
}
