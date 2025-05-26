package router

import (
	"server/internal/app"

	"github.com/gorilla/mux"
)

func SetupAuthApp(r *mux.Router) {
	auth := r.PathPrefix("/auth").Subrouter()
	
	authController := app.SetupAuthApp()

	auth.HandleFunc("/login", authController.Login).Methods("POST")
	auth.HandleFunc("/register", authController.Register).Methods("POST")
	auth.HandleFunc("/logout", authController.Logout).Methods("POST")
	auth.HandleFunc("/refresh-token", authController.RefreshToken).Methods("POST")
	auth.HandleFunc("/check", authController.CheckAuth).Methods("GET")
	auth.HandleFunc("/google", authController.LoginWithGoogle).Methods("Post")
	auth.HandleFunc("/github/login", authController.LoginWithGithub).Methods("Get")
	auth.HandleFunc("/github/callback", authController.GitHubCallback).Methods("Get")
}
