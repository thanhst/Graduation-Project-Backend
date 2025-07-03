package controller

import (
	"encoding/json"
	"net/http"
	service "server/internal/handlers/services"
	model "server/internal/models"

	"github.com/gorilla/mux"
)

type NotificationController struct {
	notificationService *service.NotificationService
}

func NewNotificationController(s *service.NotificationService) *NotificationController {
	return &NotificationController{notificationService: s}
}
func (notificationController *NotificationController) Create(w http.ResponseWriter, r *http.Request) {
	var notification model.Notification
	if err := json.NewDecoder(r.Body).Decode(&notification); err != nil {
		http.Error(w, "Fail to get notification in body", http.StatusBadRequest)
		return
	}
	if err := notificationController.notificationService.Create(&notification); err != nil {
		http.Error(w, "Error to save notification", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&notification)
}
func (notificationController *NotificationController) Update(w http.ResponseWriter, r *http.Request) {

}
func (notificationController *NotificationController) Delete(w http.ResponseWriter, r *http.Request) {
	Vars := mux.Vars(r)
	notifId := Vars["id"]
	if notifId == "" {
		http.Error(w, "Error to delete schedule", http.StatusBadRequest)
		return
	}
	if err := notificationController.notificationService.Delete(notifId); err != nil {
		http.Error(w, "Error to save notification", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Delete success!",
	})
}
func (notificationController *NotificationController) GetByClasssrom(w http.ResponseWriter, r *http.Request) {
	Vars := mux.Vars(r)
	classId := Vars["id"]
	notifications, err := notificationController.notificationService.GetByClasssrom(classId)
	if err != nil {
		http.Error(w, "Error to get notification", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&notifications)
}
func (notificationController *NotificationController) GetOne(w http.ResponseWriter, r *http.Request) {
	Vars := mux.Vars(r)
	userId := Vars["id"]
	notification, err := notificationController.notificationService.GetLatestByUserClassrooms(userId)
	if err != nil {
		http.Error(w, "Error to get notification", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&notification)
}

func (notificationController *NotificationController) GetByUserClassrooms(w http.ResponseWriter, r *http.Request) {
	Vars := mux.Vars(r)
	userId := Vars["id"]
	notifications, err := notificationController.notificationService.GetByUserClassrooms(userId)
	if err != nil {
		http.Error(w, "Error to get notification", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&notifications)
}
