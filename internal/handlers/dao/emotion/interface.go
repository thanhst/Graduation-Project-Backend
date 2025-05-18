package emotiondao

import (
	model "server/internal/models"
	"time"
)

type EmotionDAO interface {
	Create(emotion *model.Emotion) error
	GetLatestByUserInRoom(userId, roomId string) (*model.Emotion, error)
	GetAllInRoom(roomId string, fromTime time.Time) ([]model.Emotion, error)
	DeleteAllOfRoom(roomId string) error
}
