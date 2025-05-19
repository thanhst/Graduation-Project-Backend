package router

import (
	"server/internal/app"

	"github.com/gorilla/mux"
)

func SetupUserRoutes(r *mux.Router) {
	user := r.PathPrefix("/users").Subrouter()

	userController := app.Setup_User_App()

	user.HandleFunc("/", userController.CreateUser).Methods("POST")
	user.HandleFunc("/{id}", userController.GetUser).Methods("GET")
	user.HandleFunc("/{id}", userController.UpdateUser).Methods("PUT")
	user.HandleFunc("/{id}", userController.DeleteUser).Methods("DELETE")
}
