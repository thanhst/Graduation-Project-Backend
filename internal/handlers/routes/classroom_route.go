package router

import (
	"net/http"
	"server/internal/app"
	middleware "server/internal/handlers/middlewares"

	"github.com/gorilla/mux"
)

func SetupClassroomApp(r *mux.Router) {
	classroom := r.PathPrefix("/classrooms").Subrouter()
	classroom.Use(middleware.AuthMiddleware)

	studentController := app.SetupStudentClassApp()
	classController := app.SetupClassroomApp()
	// classroom.Handle("/user/{id}/alls", middleware.RoleSwitchMiddleware(
	// 	app.ServiceContainer.AccountService,
	// 	map[string]http.HandlerFunc{
	// 		"student": studentController.GetClassroomsWithUser,
	// 		"teacher": classController.GetClassroomsByUser,
	// 	},
	// )).Methods("GET")
	classroom.Handle("/user/{id}", middleware.RoleSwitchMiddleware(
		app.ServiceContainer.AccountService,
		map[string]http.HandlerFunc{
			"student": studentController.GetClassroomsWithUser,
			"teacher": classController.GetClassroomsByUser,
		},
	)).Methods("GET")
	classroom.Handle("/user/{id}/count", middleware.RoleSwitchMiddleware(
		app.ServiceContainer.AccountService,
		map[string]http.HandlerFunc{
			"student": studentController.GetCountClassroomsByUser,
			"teacher": classController.GetCountClassroomsByUser,
		},
	)).Methods("GET")
}
