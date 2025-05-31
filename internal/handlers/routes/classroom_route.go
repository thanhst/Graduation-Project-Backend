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

	classroom.HandleFunc("/{id}", classController.GetClassroomById).Methods("GET")
	classroom.HandleFunc("/{id}/count", studentController.GetCountUsersByClassroom).Methods("GET")
	classroom.HandleFunc("/{id}/users/joined", studentController.GetUserJoinedWithClassrooms).Methods("GET")
	classroom.HandleFunc("/{id}/users/waiting", studentController.GetUserWaitingWithClassrooms).Methods("GET")

	classroom.HandleFunc("/{id}/accept", middleware.RoleSwitchMiddleware(
		app.ServiceContainer.AccountService,
		map[string]http.HandlerFunc{
			"teacher": studentController.AcceptUser,
		})).Methods("POST")
	classroom.HandleFunc("/{id}/reject", middleware.RoleSwitchMiddleware(
		app.ServiceContainer.AccountService,
		map[string]http.HandlerFunc{
			"teacher": studentController.RejectUser,
		})).Methods("POST")
	classroom.HandleFunc("/{id}/join", middleware.RoleSwitchMiddleware(
		app.ServiceContainer.AccountService,
		map[string]http.HandlerFunc{
			"student": studentController.JoinClass,
		})).Methods("POST")
}
