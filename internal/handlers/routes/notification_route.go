package router

import (
	"server/internal/app"

	"github.com/gorilla/mux"
)

func SetupNotificationApp(r *mux.Router) {
	notification := r.PathPrefix("/notification").Subrouter()

	notificationController := app.SetupNotificationApp()
	notification.HandleFunc("/create", notificationController.Create).Methods("POST")
	notification.HandleFunc("/update", notificationController.Update).Methods("POST")
	notification.HandleFunc("/{id}/delete", notificationController.Delete).Methods("DELETE")
	notification.HandleFunc("/{id}/get", notificationController.GetByUserClassrooms).Methods("GET")
	notification.HandleFunc("/{id}/new", notificationController.GetOne).Methods("GET")
}
