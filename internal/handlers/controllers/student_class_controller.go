package controller

import (
	"encoding/json"
	"net/http"
	"server/internal/handlers/dto"
	service "server/internal/handlers/services"
	model "server/internal/models"
	"strconv"
	"time"

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

func (st *StudentClassController) GetUserJoinedWithClassrooms(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	classId := vars["id"]
	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")
	if classId == "" {
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
	users, err := st.stclsService.GetUserJoinedWithClassrooms(classId, limitInt, offsetInt)
	if err != nil {
		http.Error(w, "Error to get classrooms", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
func (st *StudentClassController) GetUserWaitingWithClassrooms(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	classId := vars["id"]
	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")
	if classId == "" {
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
	classrooms, err := st.stclsService.GetUserJoinedWithClassrooms(classId, limitInt, offsetInt)
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
func (st *StudentClassController) GetCountUsersByClassroom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	if userId == "" {
		http.Error(w, "Error to get userId", http.StatusBadRequest)
		return
	}
	countJoined, countWaiting, err := st.stclsService.GetCountUsersByClassroom(userId)
	if err != nil {
		http.Error(w, "Error to get classrooms", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"Count joined":  strconv.Itoa(int(countJoined)),
		"Count waiting": strconv.Itoa(int(countWaiting)),
	})
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

func (st *StudentClassController) JoinClass(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	classId := vars["id"]
	var joinRequest dto.UserReponse
	if err := json.NewDecoder(r.Body).Decode(&joinRequest); err != nil {
		http.Error(w, "Not found user in server", http.StatusInternalServerError)
		return
	}
	userId := joinRequest.UserID
	stdjoinclass := model.StudentClass{
		ClassId:   classId,
		UserId:    userId,
		State:     "waiting",
		CreatedAt: time.Now(),
		JoinedAt:  time.Now(),
	}
	std, err := st.stclsService.GetInfo(classId, userId)
	if err != nil {
		http.Error(w, "Error to accept this class", http.StatusInternalServerError)
		return
	}
	if std.UserId != "" {
		http.Error(w, "You have requested this class, please wait for teacher's response.", http.StatusInternalServerError)
		return
	}
	if err := st.stclsService.JoinClass(&stdjoinclass); err != nil {
		http.Error(w, "Error to join this class", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"messsage": "Your request has been sent to this classroom's teacher.",
	})
}
func (st *StudentClassController) AcceptUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	classId := vars["id"]
	var userId string
	if err := json.NewDecoder(r.Body).Decode(&userId); err != nil {
		http.Error(w, "Error to join this class", http.StatusInternalServerError)
		return
	}
	std, err := st.stclsService.GetInfo(classId, userId)
	if err != nil {
		http.Error(w, "Error to accept this class", http.StatusInternalServerError)
		return
	}
	std.State = "joined"
	std.JoinedAt = time.Now()
	if err := st.stclsService.Update(std); err != nil {
		http.Error(w, "Error to join this class", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Accept student success!",
	})
}
func (st *StudentClassController) RejectUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	classId := vars["id"]
	var userId string
	if err := json.NewDecoder(r.Body).Decode(&userId); err != nil {
		http.Error(w, "Error to join this class", http.StatusInternalServerError)
		return
	}
	std, err := st.stclsService.GetInfo(classId, userId)
	if err != nil {
		http.Error(w, "Error to accept this class", http.StatusInternalServerError)
		return
	}
	if err := st.stclsService.Delete(std); err != nil {
		http.Error(w, "Error to join this class", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"messsage": "Reject student success!",
	})
}
