package app

import (
	"server/internal/app/container"
	"server/internal/db/database"
	controller "server/internal/handlers/controllers"
)

var ServiceContainer *container.ServiceContainer

func Start() {
	daoContainer := container.NewDAOContainer(database.GDB)
	repoContainer := container.NewRepoContainer(daoContainer)
	ServiceContainer = container.NewServiceContainer(repoContainer)
}

func SetupUserApp() *controller.UserController {
	userController := controller.
		NewUserController(ServiceContainer.UserService,
			ServiceContainer.AccountService)
	return userController
}

func SetupAuthApp() *controller.AccountController {
	accountController := controller.
		NewAccountController(ServiceContainer.AccountService,
			ServiceContainer.UserService)
	return accountController
}

func SetupClassroomApp() *controller.ClassroomController {
	classroomController := controller.
		NewClassroomController(ServiceContainer.ClassroomService)
	return classroomController
}
func SetupStudentClassApp() *controller.StudentClassController {
	studentClassController := controller.
		NewStudentClassController(ServiceContainer.StudentClassService)
	return studentClassController
}
func SetupSchedulerApp() *controller.SchedulerController {
	schedulerController := controller.NewSchedulerController(ServiceContainer.SchedulerService, ServiceContainer.ClassroomService,
		ServiceContainer.RoomService, ServiceContainer.NotificationService)
	return schedulerController
}
func SetupRoomApp() *controller.RoomController {
	roomController := controller.NewRoomController(ServiceContainer.RoomService)
	return roomController
}
func SetupEmotionApp() *controller.EmotionController {
	emotionController := controller.NewEmotionController(ServiceContainer.EmotionService)
	return emotionController
}
func SetupNotificationApp() *controller.NotificationController {
	notificationController := controller.NewNotificationController(ServiceContainer.NotificationService)
	return notificationController
}
