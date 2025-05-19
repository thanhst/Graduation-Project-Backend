package repository

import (
	emotiondao "server/internal/handlers/dao/emotion"
	model "server/internal/models"
	"time"
)

type EmotionRepository interface {
	Create(emotion *model.Emotion) error
	GetLatestByUserInRoom(userId, roomId string) (*model.Emotion, error)
	GetAllInRoom(roomId string, fromTime time.Time) ([]model.Emotion, error)
	DeleteAllOfRoom(roomId string) error
}
type emotionRepository struct {
	emotionDAO emotiondao.EmotionDAO
}

func NewEmotionRepository(dao emotiondao.EmotionDAO) EmotionRepository {
	return &emotionRepository{emotionDAO: dao}
}

func (r *emotionRepository) Create(emotion *model.Emotion) error {
	return r.emotionDAO.Create(emotion)
}

func (r *emotionRepository) GetLatestByUserInRoom(userId, roomId string) (*model.Emotion, error) {
	return r.emotionDAO.GetLatestByUserInRoom(userId, roomId)
}

func (r *emotionRepository) GetAllInRoom(roomId string, fromTime time.Time) ([]model.Emotion, error) {
	return r.emotionDAO.GetAllInRoom(roomId, fromTime)
}

func (r *emotionRepository) DeleteAllOfRoom(roomId string) error {
	return r.emotionDAO.DeleteAllOfRoom(roomId)
}
