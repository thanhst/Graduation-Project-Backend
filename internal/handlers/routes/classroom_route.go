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
	classroom.Handle("/user/{id}/newest", middleware.RoleSwitchMiddleware(
		app.ServiceContainer.AccountService,
		map[string]http.HandlerFunc{
			"student": studentController.GetClassroomsWithNewScheduler,
			"teacher": classController.GetClassroomsWithNewScheduler,
		},
	)).Methods("GET")
	classroom.Handle("/create", middleware.RoleSwitchMiddleware(
		app.ServiceContainer.AccountService,
		map[string]http.HandlerFunc{
			"teacher": classController.Create,
		},
	)).Methods("POST")
	classroom.Handle("/{id}/update", middleware.RoleSwitchMiddleware(
		app.ServiceContainer.AccountService,
		map[string]http.HandlerFunc{
			"teacher": classController.Update,
		},
	)).Methods("PUT")
	classroom.Handle("/{id}/delete", middleware.RoleSwitchMiddleware(
		app.ServiceContainer.AccountService,
		map[string]http.HandlerFunc{
			"teacher": classController.Delete,
		},
	)).Methods("DELETE")
	classroom.HandleFunc("/{id}/teacher", classController.GetTeacherFromClass).Methods("Get")
	classroom.HandleFunc("/{id}/accept", middleware.RoleSwitchMiddleware(
		app.ServiceContainer.AccountService,
		map[string]http.HandlerFunc{
			"teacher": studentController.AcceptUser,
		}))
	classroom.HandleFunc("/{id}/reject", middleware.RoleSwitchMiddleware(
		app.ServiceContainer.AccountService,
		map[string]http.HandlerFunc{
			"teacher": studentController.RejectUser,
		}))
	classroom.HandleFunc("/{id}/join", middleware.RoleSwitchMiddleware(
		app.ServiceContainer.AccountService,
		map[string]http.HandlerFunc{
			"student": studentController.JoinClass,
		}))
}
