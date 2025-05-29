package controller

import (
	"encoding/json"
	"net/http"
	service "server/internal/handlers/services"
	model "server/internal/models"
)

type SchedulerController struct {
	schedulerService *service.SchedulerService
	classService     *service.ClassService
}

func NewSchedulerController(s *service.SchedulerService, cls *service.ClassService) *SchedulerController {
	return &SchedulerController{schedulerService: s, classService: cls}
}
func (schedulerController *SchedulerController) Create(w http.ResponseWriter, r *http.Request) {
	var scheduler model.Scheduler
	if err := json.NewDecoder(r.Body).Decode(&scheduler); err != nil {
		http.Error(w, "Fail to get schedule in body", http.StatusBadRequest)
		return
	}
	if err := schedulerController.schedulerService.Create(&scheduler); err != nil {
		http.Error(w, "Error to save schedule", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&scheduler)
}
func (schedulerController *SchedulerController) Update(w http.ResponseWriter, r *http.Request) {

}
func (schedulerController *SchedulerController) Delete(w http.ResponseWriter, r *http.Request) {

}
func (schedulerController *SchedulerController) GetSchedulerByUserAndDate(w http.ResponseWriter, r *http.Request) {

}
