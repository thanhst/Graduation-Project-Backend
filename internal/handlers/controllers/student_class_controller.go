package controller

import (
	"encoding/json"
	"net/http"
	service "server/internal/handlers/services"
	"strconv"

	"github.com/gorilla/mux"
)

type StudentClassController struct {
	stclsService *service.StudentClassService
}

func NewStudentClassController(stclsService *service.StudentClassService) *StudentClassController {
	return &StudentClassController{stclsService: stclsService}
}

//	func (st *StudentClassController) GetAllClassroomsByUser(w http.ResponseWriter, r *http.Request) {
//		vars := mux.Vars(r)
//		userId := vars["id"]
//		if userId == "" {
//			http.Error(w, "Cannot get user_id", http.StatusBadGateway)
//			return
//		}
//		classrooms, err := st.stclsService.GetAllClassroomsByUser(userId)
//		if err != nil {
//			http.Error(w, "Error to get classrooms", http.StatusBadRequest)
//			return
//		}
//		json.NewEncoder(w).Encode(classrooms)
//	}
func (st *StudentClassController) GetClassroomsWithUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")
	if userId == "" {
		http.Error(w, "Cannot get user_id", http.StatusBadGateway)
		return
	}
	if limit == "" {
		http.Error(w, "Cannot get limit", http.StatusBadGateway)
		return
	}
	if offset == "" {
		http.Error(w, "Cannot get offset", http.StatusBadGateway)
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
	classrooms, err := st.stclsService.GetClassroomsByUser(userId, limitInt, offsetInt)
	if err != nil {
		http.Error(w, "Error to get classrooms", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(classrooms)
}

func (st *StudentClassController) GetCountClassroomsByUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	if userId == "" {
		http.Error(w, "Error to get userId", http.StatusBadRequest)
		return
	}
	count, err := st.stclsService.GetCountClassroomsByUser(userId)
	if err != nil {
		http.Error(w, "Error to get classrooms", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(count)
}
func (st *StudentClassController) GetClassroomsWithNewScheduler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	classrooms, err := st.stclsService.GetClassroomsWithNewScheduler(userId)
	if err != nil {
		http.Error(w, "Error to get classrooms", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(classrooms)
}
