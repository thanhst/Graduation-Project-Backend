package controller

import (
	"encoding/json"
	"net/http"
	service "server/internal/handlers/services"
	model "server/internal/models"
	CustomHash "server/internal/utils/hash"
	"time"

	"github.com/gorilla/mux"
)

type RoomController struct {
	roomService *service.RoomService
}

func NewRoomController(r *service.RoomService) *RoomController {
	return &RoomController{roomService: r}
}

func (rm *RoomController) Create(w http.ResponseWriter, r *http.Request) {
	var room model.Room
	if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
		http.Error(w, "Error to get room!", http.StatusInternalServerError)
		return
	}
	room.RoomId = CustomHash.HashMD5(time.Now().String())
	room.CreatedAt = time.Now()
	room.State = "opening"
	if roomC, err := rm.roomService.RoomRepo.GetByID(room.RoomId); err != nil {
		http.Error(w, "Error to get room!", http.StatusInternalServerError)
		return
	} else {
		if roomC.RoomId != "" {
			if roomC.State == "opening" {
				http.Error(w, "The room is opening!!", http.StatusInternalServerError)
				return
			} else if roomC.Host != room.Host {
				http.Error(w, "Join the waiting room!!", http.StatusInternalServerError)
				return
			}
		}
	}
	if err := rm.roomService.Create(&room); err != nil {
		http.Error(w, "Error to create room!", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&room)
}

func (rm *RoomController) Update(w http.ResponseWriter, r *http.Request) {
	var room model.Room
	if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
		http.Error(w, "Error to get room!", http.StatusInternalServerError)
		return
	}
	if err := rm.roomService.RoomRepo.Update(&room); err != nil {
		http.Error(w, "Error to update room!", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&room)
}

func (rm *RoomController) Delete(w http.ResponseWriter, r *http.Request) {
	Vars := mux.Vars(r)
	roomId := Vars["id"]

	if err := rm.roomService.RoomRepo.Delete(roomId); err != nil {
		http.Error(w, "Error to update room!", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Success to delete room",
	})
}

func (rm *RoomController) GetById(w http.ResponseWriter, r *http.Request) {
	Vars := mux.Vars(r)
	roomId := Vars["id"]

	room, err := rm.roomService.RoomRepo.GetByID(roomId)
	if err != nil {
		http.Error(w, "Error to get room!", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&room)
}
