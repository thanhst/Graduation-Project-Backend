package router

import (
	"server/internal/app"

	"github.com/gorilla/mux"
)

func SetupEmotionApp(r *mux.Router) {
	emotion := r.PathPrefix("/emotion").Subrouter()

	emotionController := app.SetupEmotionApp()

	emotion.HandleFunc("/create", emotionController.Create).Methods("POST")
	emotion.HandleFunc("/update", emotionController.Update).Methods("POST")
	emotion.HandleFunc("/{id}/delete", emotionController.Delete).Methods("DELETE")
	emotion.HandleFunc("/{id}/get", emotionController.GetByRoomId).Methods("GET")
}
