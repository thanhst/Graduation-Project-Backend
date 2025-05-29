package controller

import (
	"net/http"
	service "server/internal/handlers/services"
)

type NotificationController struct {
	NotificationService *service.NotificationService
	classService        *service.ClassService
}

func NewNotificationController(s *service.NotificationService, cls *service.ClassService) *NotificationController {
	return &NotificationController{NotificationService: s, classService: cls}
}
func (notificationController *NotificationController) Create(w http.ResponseWriter, r *http.Request) {

}
func (notificationController *NotificationController) Update(w http.ResponseWriter, r *http.Request) {

}
func (notificationController *NotificationController) Delete(w http.ResponseWriter, r *http.Request) {

}
