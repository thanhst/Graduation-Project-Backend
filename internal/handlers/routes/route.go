package router

import (
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()

	user := api.PathPrefix("/users").Subrouter()
	user.HandleFunc("/", controller.CreateUser).Methods("POST")
	user.HandleFunc("/{id}", controller.GetUser).Methods("GET")
	user.HandleFunc("/{id}", controller.UpdateUser).Methods("PUT")
	user.HandleFunc("/{id}", controller.DeleteUser).Methods("DELETE")

	return r
}
