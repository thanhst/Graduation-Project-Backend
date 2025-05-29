package controller

import (
	"net/http"
	service "server/internal/handlers/services"
)

type EmotionController struct {
	EmotionService *service.EmotionService
	classService   *service.ClassService
}

func NewEmotionController(s *service.EmotionService, cls *service.ClassService) *EmotionController {
	return &EmotionController{EmotionService: s, classService: cls}
}
func (emotionController *EmotionController) Create(w http.ResponseWriter, r *http.Request) {

}
func (emotionController *EmotionController) Update(w http.ResponseWriter, r *http.Request) {

}
func (emotionController *EmotionController) Delete(w http.ResponseWriter, r *http.Request) {

}
