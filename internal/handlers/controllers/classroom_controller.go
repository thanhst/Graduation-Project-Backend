package controller

import service "server/internal/handlers/services"

type ClassroomController struct {
	classservice *service.ClassService
}

func NewClassroomController(a *service.AccountService, u *service.UserService) *AccountController {
	return &AccountController{accountService: a, userService: u}
}
