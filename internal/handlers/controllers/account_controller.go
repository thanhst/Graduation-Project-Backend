package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/internal/handlers/dto"
	service "server/internal/handlers/services"
	model "server/internal/models"
	"server/internal/utils/dotenv"
	CustomHash "server/internal/utils/hash"
	"server/internal/utils/helper"
	jwtutil "server/internal/utils/jwt"
	"time"

	"cloud.google.com/go/auth/credentials/idtoken"
)

type AccountController struct {
	accountService *service.AccountService
	userService    *service.UserService
}

func NewAccountController(a *service.AccountService, u *service.UserService) *AccountController {
	return &AccountController{accountService: a, userService: u}
}

func (ac *AccountController) Register(w http.ResponseWriter, r *http.Request) {
	var registerRequest dto.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&registerRequest); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	accounts, _ := ac.accountService.GetFullAccountWithEmail(registerRequest.Account.Email)
	if len(accounts) > 0 {
		foundEmailLogin := false
		existingUserId := accounts[0].User.UserId

		for _, acc := range accounts {
			if acc.LoginMethod == "email" {
				foundEmailLogin = true
				break
			}
		}

		if foundEmailLogin {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"error":   "email already registered with email login method",
				"message": helper.CapitalizeFirstLetter("email already registered with email login method."),
			})
			return
		}
		registerRequest.Account.UserId = existingUserId
		err := ac.accountService.Register(&registerRequest.Account)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"error":   err.Error(),
				"message": helper.CapitalizeFirstLetter(err.Error()),
			})
			return
		}
	} else {
		user := model.User{
			UserId:         CustomHash.HashMD5(time.Now().String()),
			FullName:       registerRequest.User.FullName,
			ProfilePicture: helper.RandomeImagesURL(),
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}
		if err := ac.userService.CreateUser(&user); err == nil {
			registerRequest.Account.UserId = user.UserId
			err := ac.accountService.Register(&registerRequest.Account)
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{
					"error":   err.Error(),
					"message": helper.CapitalizeFirstLetter(err.Error()),
				})
				return
			}
		}
	}
	log.Printf("%v register!\n", registerRequest.Account.UserId)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User registered successfully",
	})
}

func (ac *AccountController) Login(w http.ResponseWriter, r *http.Request) {
	var input model.Account
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	tokens, err := ac.accountService.Login(&input)
	if err != nil {
		http.Error(w, helper.CapitalizeFirstLetter(err.Error()), http.StatusUnauthorized)
		return
	}
	helper.SetTokenCookies(w, tokens["access_token"], tokens["refresh_token"], tokens["access_exprise"], tokens["refresh_exprise"])
	log.Printf("%v login!\n", tokens["user_id"])

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"user_id":         tokens["user_id"],
		"role":            tokens["role"],
		"last_login":      tokens["last_login"],
		"access_exprise":  tokens["access_exprise"],
		"refresh_exprise": tokens["refresh_exprise"],
	}
	json.NewEncoder(w).Encode(response)
}

func (ac *AccountController) Logout(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	var userId string
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	userId = data["userId"].(string)
	if userId != "" {
		err := ac.accountService.Logout(&userId)
		if err != nil {
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Opps! Server have error!",
			})
		}
		helper.RemoveCookies(w)
		log.Printf("%v logout!\n", userId)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Logout successfully!",
		})
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Cannot found the user id!",
		})
	}
}

func (ac *AccountController) RefreshToken(w http.ResponseWriter, r *http.Request) {
	refreshTokenCookie, err := r.Cookie("refresh_token")
	if err != nil {
		http.Error(w, "Missing refresh token", http.StatusUnauthorized)
		return
	}

	tokens, err := jwtutil.RefreshTokens(refreshTokenCookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	helper.SetTokenCookies(w, tokens["access_token"], tokens["refresh_token"], tokens["access_exprise"], tokens["refresh_exprise"])

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"access_exprise":  tokens["access_exprise"],
		"refresh_exprise": tokens["refresh_exprise"],
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (ac *AccountController) CheckAuth(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access_token")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	tokenStr := cookie.Value

	userId, err := jwtutil.VerifyAccessToken(tokenStr)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"userId":  userId,
		"message": "User is logged in",
	})
}

func (ac *AccountController) LoginWithGoogle(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Token string `json:"token"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	ctx := context.Background()

	validator, err := idtoken.NewValidator(nil)
	if err != nil {
		http.Error(w, "Failed to create validator", http.StatusInternalServerError)
		return
	}

	payload, err := validator.Validate(ctx, req.Token, dotenv.GetDotEnv("GOOGLE_CLIENT_ID"))
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	email, ok := payload.Claims["email"].(string)
	if !ok {
		http.Error(w, "Email claim missing", http.StatusUnauthorized)
		return
	}
	name, ok := payload.Claims["name"].(string)
	if !ok {
		name = "UserGoogle"
	}
	//check account by email. And create user if necessary
	var user *model.User
	if accounts, _ := ac.accountService.GetFullAccountWithEmail(email); len(accounts) == 0 {
		user = &model.User{
			UserId:         CustomHash.HashMD5(time.Now().String()),
			FullName:       name,
			ProfilePicture: helper.RandomeImagesURL(),
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}
		ac.userService.CreateUser(user)
	} else {
		user = &accounts[0].User
	}
	Account := &model.Account{
		UserId:      user.UserId,
		Email:       email,
		LoginMethod: "google",
	}

	if tokens, err := ac.accountService.LoginWithGoogle(Account); err == nil {
		helper.SetTokenCookies(w, tokens["access_token"], tokens["refresh_token"], tokens["access_exprise"], tokens["refresh_exprise"])
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"user_id":         tokens["user_id"],
			"role":            tokens["role"],
			"last_login":      tokens["last_login"],
			"access_exprise":  tokens["access_exprise"],
			"refresh_exprise": tokens["refresh_exprise"],
		}
		log.Println(tokens["user_id"] + " login!")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
}

func (a *AccountController) LoginWithGithub(w http.ResponseWriter, r *http.Request) {
	reUri := r.URL.Query().Get("redirect_uri")
	state := dto.GitHubUserRequest{
		RedirectURI: reUri,
		Nonce:       CustomHash.HashMD5(time.Now().String()),
		ExpiresAt:   time.Now().Add(5 * time.Minute).Unix(),
	}
	url := a.accountService.LoginWithGithub(&state)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (a *AccountController) GitHubCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")
	if code == "" {
		http.Error(w, "Missing code in request", http.StatusBadRequest)
		return
	}
	if state == "" {
		http.Error(w, "Missing state in request", http.StatusBadRequest)
	}
	claims, err := jwtutil.VerifyAccessTokenGithub(state)
	if err != nil {
		http.Error(w, "Error to convert claims", http.StatusInternalServerError)
	}
	jsonClaims, err := json.Marshal(claims)
	if err != nil {
		http.Error(w, "Error to convert claims", http.StatusInternalServerError)
		return
	}
	//convert to dto
	var stateDto dto.GitHubUserRequest
	if err := json.Unmarshal(jsonClaims, &stateDto); err != nil {
		http.Error(w, "Error to convert claims to dto", http.StatusInternalServerError)
		return
	}

	user, err := a.accountService.GitHubCallback(code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if accounts, err := a.accountService.GetFullAccountWithEmail(user.Email); err == nil {
		now := time.Now()
		account, _ := a.accountService.GetAccountByEmailAndMethod(user.Email, "github")
		if account.AccountId != "" {
			account.Status = "online"
			if !a.accountService.Update(account) {
				http.Error(w, "Error for create account github", http.StatusInternalServerError)
				return
			}
		} else {
			account.AccountId = CustomHash.HashMD5(time.Now().String())
			account.Email = user.Email
			account.CreatedAt = time.Now()
			account.UpdatedAt = time.Now()
			account.LoginMethod = "github"
			account.LastLogin = &now
			if len(accounts) > 0 {
				account.UserId = accounts[0].UserId
				account.Role = accounts[0].Role
			} else {
				userDB := model.User{
					UserId:         CustomHash.HashMD5(time.Now().String()),
					FullName:       "",
					ProfilePicture: helper.RandomeImagesURL(),
					CreatedAt:      time.Now(),
					UpdatedAt:      time.Now(),
				}
				a.userService.CreateUser(&userDB)
				account.UserId = userDB.UserId
				account.Role = "student"
			}
			if err := a.accountService.Create(account); err != nil {
				http.Error(w, "Error for create account github", http.StatusInternalServerError)
				return
			}
		}

		accessToken, err := jwtutil.CreateToken(account.UserId, jwtutil.GetAccessToken(), jwtutil.GetAccessExpire())
		if err != nil {
			http.Error(w, "Error to convert claims to dto", http.StatusInternalServerError)
			return
		}
		refreshToken, err := jwtutil.CreateToken(account.UserId, jwtutil.GetRefreshToken(), jwtutil.GetRefreshExpire())
		if err != nil {
			http.Error(w, "Error to convert claims to dto", http.StatusInternalServerError)
			return
		}
		var last_login string
		LastLogin := account.LastLogin
		if LastLogin == nil {
			last_login = ""
		} else {
			last_login = LastLogin.String()
		}

		helper.SetTokenCookies(w, accessToken, refreshToken, time.Now().Add(jwtutil.GetAccessExpire()).String(), time.Now().Add(jwtutil.GetRefreshExpire()).String())
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"user_id":         account.UserId,
			"role":            account.Role,
			"last_login":      last_login,
			"access_exprise":  time.Now().Add(jwtutil.GetAccessExpire()).String(),
			"refresh_exprise": time.Now().Add(jwtutil.GetRefreshExpire()).String(),
		}
		jsonBytes, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Error to convert claims to dto", http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		html := fmt.Sprintf(`<!DOCTYPE html>
			<html>
			<head><title>OAuth Callback</title></head>
			<body>
			<script>
			const data = %s;
			window.opener.postMessage(data, "%s");
			window.close();
			</script>
			</body>
			</html>`, string(jsonBytes), stateDto.RedirectURI)
		log.Println(account.UserId + " login!")
		w.Write([]byte(html))
	} else {
		http.Error(w, "Error of internal server!", http.StatusInternalServerError)
	}
}
