package controller

import (
	"context"
	"encoding/json"
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

func (ac *AccountController) GoogleLoginHandler(w http.ResponseWriter, r *http.Request) {
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
	// name, ok := payload.Claims["name"].(string)
	// if !ok {
	// 	name = "UserGoogle"
	// }
	//check account by email. And create user if necessary
	jwtToken, err := jwtutil.CreateToken(email, jwtutil.GetAccessToken(), jwtutil.GetAccessExpire())
	refreshToken, err := jwtutil.CreateToken(email, jwtutil.GetRefreshToken(), jwtutil.GetRefreshExpire())

	helper.SetTokenCookies(w, jwtToken, refreshToken, jwtutil.GetAccessExpire().String(), jwtutil.GetRefreshExpire().String())

	json.NewEncoder(w).Encode(map[string]string{})
}
