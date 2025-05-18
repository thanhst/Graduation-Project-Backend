package emotiondao

import (
	model "server/internal/models"
	"time"

	"gorm.io/gorm"
)

type emotionDAOImpl struct {
	db *gorm.DB
}

func NewEmotionDAO(db *gorm.DB) EmotionDAO {
	return &emotionDAOImpl{db: db}
}

func (dao *emotionDAOImpl) Create(emotion *model.Emotion) error {
	return dao.db.Create(emotion).Error
}

func (dao *emotionDAOImpl) GetLatestByUserInRoom(userId, roomId string) (*model.Emotion, error) {
	var emo model.Emotion
	err := dao.db.
		Where("user_id = ? AND room_id = ?", userId, roomId).
		Order("created_at DESC").
		First(&emo).Error
	return &emo, err
}

func (dao *emotionDAOImpl) GetAllInRoom(roomId string, fromTime time.Time) ([]model.Emotion, error) {
	var emos []model.Emotion
	err := dao.db.
		Where("room_id = ? AND created_at >= ?", roomId, fromTime).
		Order("created_at ASC").
		Find(&emos).Error
	return emos, err
}

func (dao *emotionDAOImpl) DeleteAllOfRoom(roomId string) error {
	return dao.db.Where("room_id = ?", roomId).Delete(&model.Emotion{}).Error
}
