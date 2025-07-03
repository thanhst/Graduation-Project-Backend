package controller

import (
	"encoding/json"
	"net/http"
	service "server/internal/handlers/services"
	model "server/internal/models"

	"github.com/gorilla/mux"
)

type EmotionController struct {
	EmotionService *service.EmotionService
}

func NewEmotionController(s *service.EmotionService) *EmotionController {
	return &EmotionController{EmotionService: s}
}
func (emotionController *EmotionController) Create(w http.ResponseWriter, r *http.Request) {
	var emotion model.Emotion
	if err := json.NewDecoder(r.Body).Decode(&emotion); err != nil {
		http.Error(w, "Fail to get emotion in body", http.StatusBadRequest)
		return
	}
	if err := emotionController.EmotionService.Create(&emotion); err != nil {
		http.Error(w, "Error to save Emotion", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&emotion)
}
func (emotionController *EmotionController) Update(w http.ResponseWriter, r *http.Request) {

}
func (emotionController *EmotionController) Delete(w http.ResponseWriter, r *http.Request) {

}
func (emotionController *EmotionController) GetByUserId(w http.ResponseWriter, r *http.Request) {

}
func (emotionController *EmotionController) GetByRoomId(w http.ResponseWriter, r *http.Request) {
	Vars := mux.Vars(r)
	roomId := Vars["id"]
	if roomId != "" {
		emotions, err := emotionController.EmotionService.GetByRoomId(roomId)
		if err != nil {
			http.Error(w, "Error to get Emotion", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(&emotions)
	} else {
		http.Error(w, "Error to get id from this room", http.StatusBadRequest)
		return
	}
}
