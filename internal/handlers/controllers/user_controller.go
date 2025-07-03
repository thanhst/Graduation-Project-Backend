package controller

import (
	"encoding/json"
	"io"
	"log"
	"math/rand/v2"
	"net/http"
	"os"
	"path/filepath"
	service "server/internal/handlers/services"
	model "server/internal/models"
	"server/internal/utils/dotenv"
	CustomHash "server/internal/utils/hash"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type UserController struct {
	userService    *service.UserService
	accountService *service.AccountService
}

func NewUserController(userService *service.UserService, accountService *service.AccountService) *UserController {
	return &UserController{userService: userService, accountService: accountService}
}

func (userController *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := userController.userService.CreateUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (userController *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	// log.Println("Get user")
	id := mux.Vars(r)["id"]
	user, err := userController.userService.GetUserByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (userController *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(50 << 20)
	if err != nil {
		http.Error(w, "Cannot parse multipart form", http.StatusBadRequest)
		return
	}
	var user model.User
	role := r.FormValue("selectedRole")
	teacherId := r.FormValue("teacherId")
	userData := r.FormValue("user")
	if userData == "" {
		http.Error(w, "Missing user data", http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal([]byte(userData), &user); err != nil {
		http.Error(w, "Invalid user JSON data", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("imagePreview")

	if err == nil {
		files, err := filepath.Glob("./uploads/" + user.UserId + ".*")
		if err == nil {
			for _, f := range files {
				os.Remove(f)
			}
		}
		defer file.Close()
		ext := filepath.Ext(header.Filename)
		filePath := "/uploads/" + user.UserId + ext
		dst, err := os.Create("." + filePath)
		if err != nil {
			http.Error(w, "Cannot save image"+err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		_, err = io.Copy(dst, file)
		if err != nil {
			http.Error(w, "Error to save image!", http.StatusInternalServerError)
			return
		}
		if dotenv.GetDotEnv("APP_ENV") == "production" {
			user.ProfilePicture = dotenv.GetDotEnv("APP_URL") + filePath
		} else {
			user.ProfilePicture = dotenv.GetDotEnv("APP_URL") + ":" + dotenv.GetDotEnv("APP_PORT") + filePath
		}
	} else {
		files, err := filepath.Glob("./uploads/" + user.UserId + ".*")
		if err == nil {
			for _, f := range files {
				os.Remove(f)
			}
		}
		if dotenv.GetDotEnv("APP_ENV") == "production" {
			user.ProfilePicture = dotenv.GetDotEnv("APP_URL") + "/uploads/" + strconv.Itoa(rand.IntN(8)+1) + ".jpg"
		} else {
			user.ProfilePicture = dotenv.GetDotEnv("APP_URL") + ":" + dotenv.GetDotEnv("APP_PORT") + "/uploads/" + strconv.Itoa(rand.IntN(8)+1) + ".jpg"
		}
		// http.Error(w, "Cannot get image: "+err.Error(), http.StatusBadRequest)
		// return
	}
	if accounts, err := userController.accountService.GetFullAccountWithUser(user.UserId); err == nil {
		if role != "" {
			for i := range accounts {
				accounts[i].Role = role
				now := time.Now()
				accounts[i].LastLogin = &now
				if !userController.accountService.Update(accounts[i]) {
					http.Error(w, "Error fo update role!", http.StatusInternalServerError)
					return
				}
			}
		} else {
			log.Printf("Not set role for user")
		}
		if teacherId != "" {
		}

		if err := userController.userService.UpdateUser(&user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"user":       user,
			"role":       role,
			"last_login": time.Now().String(),
		})
	} else {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Error for update user",
		})
	}
}

func (userController *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := userController.userService.DeleteUser(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
func (userController *UserController) UpdateInformationOfUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := r.ParseMultipartForm(50 << 20)
	if err != nil {
		http.Error(w, "Cannot parse form", http.StatusBadRequest)
		return
	}
	fullname := r.FormValue("fullname")
	oldPassword := r.FormValue("oldPassword")
	newPassword := r.FormValue("newPassword")
	loginMethod := r.FormValue("loginMethod")

	if oldPassword != "" && newPassword != "" {
		accounts, err := userController.accountService.GetAccountsByUser(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			if len(accounts) > 0 {
				if loginMethod == "email" {
					for i := 0; i < len(accounts); i++ {
						if accounts[i].LoginMethod == "email" {
							if CustomHash.CheckPassword(accounts[i].Password, oldPassword) {
								newPass, err := CustomHash.HashPassword(newPassword)
								if err != nil {
									http.Error(w, err.Error(), http.StatusInternalServerError)
									return
								}
								accounts[i].Password = newPass
								userController.accountService.Update(accounts[i])
							} else {
								http.Error(w, "Your password is incorrect!", http.StatusInternalServerError)
								return
							}
							break
						}
					}
				}
			} else {
				http.Error(w, "Error for find accounts", http.StatusInternalServerError)
				return
			}
		}
	}

	existingUser, err := userController.userService.GetUserByID(id)
	if err != nil {
		http.Error(w, "Error for find your account!", http.StatusInternalServerError)
		return
	} else {
		file, header, err := r.FormFile("imagePreview")

		if err == nil {
			files, err := filepath.Glob("./uploads/" + id + ".*")
			if err == nil {
				for _, f := range files {
					os.Remove(f)
				}
			}
			defer file.Close()
			ext := filepath.Ext(header.Filename)
			filePath := "/uploads/" + id + ext
			dst, err := os.Create("." + filePath)
			if err != nil {
				http.Error(w, "Cannot save image"+err.Error(), http.StatusInternalServerError)
				return
			}
			defer dst.Close()

			_, err = io.Copy(dst, file)
			if err != nil {
				http.Error(w, "Error to save image!", http.StatusInternalServerError)
				return
			}
			if dotenv.GetDotEnv("APP_ENV") == "production" {
				existingUser.ProfilePicture = dotenv.GetDotEnv("APP_URL") + filePath
			} else {
				existingUser.ProfilePicture = dotenv.GetDotEnv("APP_URL") + ":" + dotenv.GetDotEnv("APP_PORT") + filePath
			}
		}
		existingUser.FullName = fullname
	}

	errOfupdate := userController.userService.UpdateUser(existingUser)
	if errOfupdate != nil {
		http.Error(w, "Error for update your account!", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"user": existingUser})
}
