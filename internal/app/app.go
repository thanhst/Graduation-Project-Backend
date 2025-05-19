package app

import (
	"server/internal/db/database"
	controller "server/internal/handlers/controllers"
	userdao "server/internal/handlers/dao/user"
	service "server/internal/handlers/services"
)

func Setup_User_App() *controller.UserController {
	userDAO := userdao.NewUserDAO(database.GDB)
	userService := service.NewUserService(userDAO)
	userController := controller.NewUserController(userService)
	return userController
}
