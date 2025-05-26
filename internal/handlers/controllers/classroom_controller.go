package controller

import (
	"encoding/json"
	"net/http"
	service "server/internal/handlers/services"
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
func (cls *ClassroomController) UpdateClass(w http.ResponseWriter, r *http.Request) {}
func (cls *ClassroomController) DeleteClass(w http.ResponseWriter, r *http.Request) {}
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
func (cls *ClassroomController) GetClassroomById(w http.ResponseWriter, r *http.Request) {}
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
