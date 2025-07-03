package service

import (
	repository "server/internal/handlers/repositories"
	model "server/internal/models"
)

type EmotionService struct {
	EmotionRepo repository.EmotionRepository
}

func NewEmotionService(EmotionRepo repository.EmotionRepository) *EmotionService {
	return &EmotionService{EmotionRepo: EmotionRepo}
}
func (emotionService *EmotionService) Create(schedule *model.Emotion) error {
	return emotionService.EmotionRepo.Create(schedule)
}
func (emotionService *EmotionService) GetByRoomId(roomId string) ([]*model.Emotion, error) {
	return emotionService.EmotionRepo.GetByRoomId(roomId)
}

// func (emotionService *EmotionService) Update(schedule *model.Emotion) error {
// 	return emotionService.EmotionRepo.Update(schedule)
// }
// func (emotionService *EmotionService) Delete(scheduleId string) error {
// 	return emotionService.EmotionRepo.Delete(scheduleId)
// }
