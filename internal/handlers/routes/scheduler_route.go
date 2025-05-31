package router

import (
	"server/internal/app"
	middleware "server/internal/handlers/middlewares"

	"github.com/gorilla/mux"
)

func SetupSchedulerRoutes(r *mux.Router) {
	scheduler := r.PathPrefix("/schedulers").Subrouter()
	scheduler.Use(middleware.AuthMiddleware)

	schedulerController := app.SetupSchedulerApp()

	scheduler.HandleFunc("/create", schedulerController.Create).Methods("POST")
	scheduler.HandleFunc("/user/{id}/get", schedulerController.GetSchedulerByUserAndDate).Methods("GET")
	scheduler.HandleFunc("/user/{id}/get/all", schedulerController.GetSchedulerByUser).Methods("GET")
	scheduler.HandleFunc("/{id}/update", schedulerController.Update).Methods("POST")
	scheduler.HandleFunc("/{id}/delete", schedulerController.Delete).Methods("DELETE")
	scheduler.HandleFunc("/{id}/view", schedulerController.View).Methods("GET")
}
