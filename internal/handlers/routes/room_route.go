package router

import (
	"server/internal/app"

	"github.com/gorilla/mux"
)

func SetupRoomApp(r *mux.Router) {
	room := r.PathPrefix("/room").Subrouter()

	roomController := app.SetupRoomApp()

	room.HandleFunc("/create", roomController.Create).Methods("POST")
	room.HandleFunc("/update", roomController.Update).Methods("POST")
	room.HandleFunc("/{id}/delete", roomController.Delete).Methods("DELETE")
	room.HandleFunc("/{id}/get", roomController.GetById).Methods("GET")
}
