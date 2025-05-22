package router

import (
	"server/internal/app"
	middleware "server/internal/handlers/middlewares"

	"github.com/gorilla/mux"
)

func SetupUserRoutes(r *mux.Router) {
	user := r.PathPrefix("/users").Subrouter()
	user.Use(middleware.AuthMiddleware)

	userController := app.SetupUserApp()

	user.HandleFunc("/create", userController.CreateUser).Methods("POST")
	user.HandleFunc("/{id}/get", userController.GetUser).Methods("GET")
	user.HandleFunc("/{id}/update", userController.UpdateUser).Methods("POST")
	user.HandleFunc("/{id}/delete", userController.DeleteUser).Methods("DELETE")
}
