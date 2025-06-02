package controller

import (
	"encoding/json"
	"log"
	"net/http"
	service "server/internal/handlers/services"
	model "server/internal/models"
	CustomHash "server/internal/utils/hash"
	"time"

	"github.com/gorilla/mux"
)

type SchedulerController struct {
	schedulerService    *service.SchedulerService
	classService        *service.ClassService
	roomService         *service.RoomService
	notificationService *service.NotificationService
}

func NewSchedulerController(s *service.SchedulerService, cls *service.ClassService,
	roomService *service.RoomService, notificationService *service.NotificationService) *SchedulerController {
	return &SchedulerController{schedulerService: s, classService: cls,
		roomService: roomService, notificationService: notificationService}
}
func (schedulerController *SchedulerController) Create(w http.ResponseWriter, r *http.Request) {
	var scheduler model.Scheduler
	if err := json.NewDecoder(r.Body).Decode(&scheduler); err != nil {
		log.Println(err)
		http.Error(w, "Fail to get schedule in body", http.StatusBadRequest)
		return
	}
	roomId := CustomHash.HashMD5(time.Now().String())
	scheduler.RoomId = roomId
	scheduler.SchedulerId = CustomHash.HashMD5(time.Now().String())
	if scheduler.ClassId != nil {
		count, err := schedulerController.schedulerService.GetCountSchedulerWithTime(*scheduler.ClassId, scheduler.StartTime)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error to check schedule", http.StatusBadRequest)
			return
		}
		if count > 0 {
			http.Error(w, "You had a scheduler! Cannot create new!", http.StatusBadRequest)
			return
		}
	}
	room := model.Room{
		RoomId:    roomId,
		ClassId:   scheduler.ClassId,
		State:     "closed",
		Host:      scheduler.UserId,
		EndedAt:   nil,
		CreatedAt: time.Now(),
	}
	if err := schedulerController.roomService.Create(&room); err != nil {
		log.Println(err)
		http.Error(w, "Error to create room", http.StatusBadRequest)
		return
	}
	var notification *model.Notification
	if scheduler.ClassId != nil {
		notification = &model.Notification{
			NotificationId: CustomHash.HashMD5(time.Now().String()),
			UserId:         scheduler.UserId,
			ClassId:        scheduler.ClassId,
			Description:    scheduler.Description,
			Type:           "success",
			CreatedAt:      time.Now(),
		}
	} else {
		notification = &model.Notification{
			NotificationId: CustomHash.HashMD5(time.Now().String()),
			UserId:         scheduler.UserId,
			Description:    scheduler.Description,
			Type:           "success",
			CreatedAt:      time.Now(),
		}
	}

	if err := schedulerController.notificationService.Create(notification); err != nil {
		log.Println(err)
		http.Error(w, "Error to save schedule, notification cannot create", http.StatusBadRequest)
		return
	}

	if err := schedulerController.schedulerService.Create(&scheduler); err != nil {
		log.Println(err)
		http.Error(w, "Error to save schedule", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&scheduler)
}
func (schedulerController *SchedulerController) Update(w http.ResponseWriter, r *http.Request) {
	var scheduler model.Scheduler
	notification := &model.Notification{
		NotificationId: CustomHash.HashMD5(time.Now().String()),
		UserId:         scheduler.UserId,
		ClassId:        scheduler.ClassId,
		Description:    scheduler.Description,
		Type:           "info",
		CreatedAt:      time.Now(),
	}
	if err := schedulerController.notificationService.Create(notification); err != nil {
		log.Println(err)
		http.Error(w, "Error to save schedule, notification cannot create", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&scheduler); err != nil {
		http.Error(w, "Fail to get schedule in body", http.StatusBadRequest)
		return
	}
	if err := schedulerController.schedulerService.Update(&scheduler); err != nil {
		http.Error(w, "Error to update schedule", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&scheduler)
}
func (schedulerController *SchedulerController) Delete(w http.ResponseWriter, r *http.Request) {
	Vars := mux.Vars(r)
	schedulerId := Vars["id"]
	if schedulerId == "" {
		http.Error(w, "Error to delete schedule", http.StatusBadRequest)
		return
	}
	if s, err := schedulerController.schedulerService.View(schedulerId); err != nil {
		log.Println(err)
		http.Error(w, "Not found the schedule in server", http.StatusBadRequest)
		return
	} else {
		notification := &model.Notification{
			NotificationId: CustomHash.HashMD5(time.Now().String()),
			UserId:         s.UserId,
			ClassId:        s.ClassId,
			Description:    s.Description,
			Type:           "info",
			CreatedAt:      time.Now(),
		}
		if err := schedulerController.notificationService.Create(notification); err != nil {
			log.Println(err)
			http.Error(w, "Error to save schedule, notification cannot create", http.StatusBadRequest)
			return
		}
	}
	if err := schedulerController.schedulerService.Delete(schedulerId); err != nil {
		http.Error(w, "Error to delete schedule", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Success to delete scheduler",
	})
}
func (schedulerController *SchedulerController) GetSchedulerByUserAndDate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	if userId == "" {
		http.Error(w, "Error to get userID", http.StatusInternalServerError)
		return
	}
	date := r.URL.Query().Get("date")
	if date == "" {
		http.Error(w, "Missing date parameter", http.StatusBadRequest)
		return
	}
	dateConvert, err := time.Parse("2006-01-02", date)
	if err != nil {
		http.Error(w, "Error to get date", http.StatusInternalServerError)
		return
	}
	schedulers, err := schedulerController.schedulerService.GetSchedulerByUserAndDate(userId, dateConvert.String())
	if err != nil {
		http.Error(w, "Error to get date", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&schedulers)
}
func (schedulerController *SchedulerController) GetSchedulerByUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	if userId == "" {
		http.Error(w, "Error to get userID", http.StatusInternalServerError)
		return
	}
	schedulers, err := schedulerController.schedulerService.GetSchedulerByUser(userId)
	if err != nil {
		http.Error(w, "Error to get date", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&schedulers)
}
func (schedulerController *SchedulerController) View(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sId := vars["id"]
	if sId == "" {
		http.Error(w, "Error to get userID", http.StatusInternalServerError)
		return
	}
	s, err := schedulerController.schedulerService.View(sId)
	if err != nil {
		http.Error(w, "Error to get scheduler", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&s)
}
