package router

import (
	"net/http"
	"server/internal/app"
	middleware "server/internal/handlers/middlewares"

	"github.com/gorilla/mux"
)

func SetupAuthApp(r *mux.Router) {
	auth := r.PathPrefix("/auth").Subrouter()

	authController := app.SetupAuthApp()

	auth.HandleFunc("/login", authController.Login).Methods("POST")
	auth.HandleFunc("/register", authController.Register).Methods("POST")
	auth.Handle("/logout", middleware.AuthMiddleware(http.HandlerFunc(authController.Logout))).Methods("POST")
	auth.HandleFunc("/refresh-token", authController.RefreshToken).Methods("POST")
	auth.HandleFunc("/check", authController.CheckAuth).Methods("GET")
	auth.HandleFunc("/google", authController.LoginWithGoogle).Methods("Post")
	auth.HandleFunc("/github/login", authController.LoginWithGithub).Methods("Get")
	auth.HandleFunc("/github/callback", authController.GitHubCallback).Methods("Get")
}
