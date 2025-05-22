package app

import (
	"server/internal/app/container"
	"server/internal/db/database"
	controller "server/internal/handlers/controllers"
)

var serviceContainer *container.ServiceContainer

func Start() {
	daoContainer := container.NewDAOContainer(database.GDB)
	repoContainer := container.NewRepoContainer(daoContainer)
	serviceContainer = container.NewServiceContainer(repoContainer)
}

func SetupUserApp() *controller.UserController {
	userController := controller.
		NewUserController(serviceContainer.UserService,
			serviceContainer.AccountService)
	return userController
}

func SetupAuthApp() *controller.AccountController {
	accountController := controller.
		NewAccountController(serviceContainer.AccountService,
			serviceContainer.UserService)
	return accountController
}
