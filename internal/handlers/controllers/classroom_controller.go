package controller

import (
	"encoding/json"
	"net/http"
	service "server/internal/handlers/services"
	model "server/internal/models"
	"strconv"

	"github.com/gorilla/mux"
)

type ClassroomController struct {
	classService *service.ClassService
}

func NewClassroomController(cls *service.ClassService) *ClassroomController {
	return &ClassroomController{classService: cls}
}

func (cls *ClassroomController) GetAllClassroomsByUser(w http.ResponseWriter, r *http.Request) {

}
func (cls *ClassroomController) Create(w http.ResponseWriter, r *http.Request) {
	var classroom model.Classroom
	if err := json.NewDecoder(r.Body).Decode(&classroom); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	err := cls.classService.Create(&classroom)
	if err != nil {
		http.Error(w, "Error to create classroom", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(classroom)
}

func (cls *ClassroomController) Update(w http.ResponseWriter, r *http.Request) {
	var classroom model.Classroom
	if err := json.NewDecoder(r.Body).Decode(&classroom); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	err := cls.classService.Update(&classroom)
	if err != nil {
		http.Error(w, "Error to update classroom", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(classroom)
}
func (cls *ClassroomController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	classId := vars["id"]
	err := cls.classService.Delete(classId)
	if err != nil {
		http.Error(w, "Error to delete classroom", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Success to delete classroom with Id : " + classId,
	})
}
func (cls *ClassroomController) GetClassroomsByUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")
	if userId == "" {
		http.Error(w, "Missing userId", http.StatusBadRequest)
		return
	}
	if limit == "" {
		http.Error(w, "Missing limit", http.StatusBadRequest)
		return
	}
	if offset == "" {
		http.Error(w, "Missing offset", http.StatusBadRequest)
		return
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		http.Error(w, "Error to convert limit int", http.StatusBadRequest)
		return
	}
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		http.Error(w, "Error to convert offset int", http.StatusBadRequest)
		return
	}
	classrooms, err := cls.classService.GetClassroomsByUser(userId, limitInt, offsetInt)
	if err != nil {
		http.Error(w, "Error to get classrooms", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(classrooms)
}
func (cls *ClassroomController) GetClassroomById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	classId := vars["id"]
	if classId == "" {
		http.Error(w, "Error to get classrooms", http.StatusBadRequest)
		return
	}
	classroom, err := cls.classService.GetClassroomById(classId)
	if err != nil {
		http.Error(w, "Error to get classrooms", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(classroom)
}
func (cls *ClassroomController) GetCountClassroomsByUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	if userId == "" {
		http.Error(w, "Error to get classrooms", http.StatusBadRequest)
		return
	}
	count, err := cls.classService.GetCountClassroomsByUser(userId)
	if err != nil {
		http.Error(w, "Error to get classrooms", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(count)
}
func (cls *ClassroomController) GetClassroomsWithNewScheduler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	classrooms, err := cls.classService.GetClassroomsWithNewScheduler(userId)
	if err != nil {
		http.Error(w, "Error to get classrooms", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(classrooms)
}
