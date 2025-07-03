package roomdao

import (
	model "server/internal/models"
	"time"

	"gorm.io/gorm"
)

type roomDAOImpl struct {
	db *gorm.DB
}

func NewRoomDAO(db *gorm.DB) RoomDAO {
	return &roomDAOImpl{db: db}
}

func (dao *roomDAOImpl) GetByID(roomID string) (*model.Room, error) {
	var room model.Room
	err := dao.db.
		Preload("Classroom").
		Preload("User").
		Where("room_id = ?", roomID).
		Find(&room).Error
	return &room, err
}

func (dao *roomDAOImpl) GetByHost(userID string, limit int, offset int) ([]model.Room, error) {
	var rooms []model.Room
	err := dao.db.
		Where("host = ? AND class_id IS NOT NULL and class_id !=''", userID).
		Preload("Classroom").
		Preload("User").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&rooms).Error
	return rooms, err
}

func (dao *roomDAOImpl) GetActiveRooms() ([]model.Room, error) {
	var rooms []model.Room
	err := dao.db.
		Where("state = ?", "opening").
		Order("created_at DESC").
		Find(&rooms).Error
	return rooms, err
}

func (dao *roomDAOImpl) Create(room *model.Room) error {
	return dao.db.Create(room).Error
}

func (dao *roomDAOImpl) Update(room *model.Room) error {
	return dao.db.Save(room).Error
}

func (dao *roomDAOImpl) CloseRoom(roomID string) error {
	now := time.Now()
	return dao.db.Model(&model.Room{}).
		Where("room_id = ?", roomID).
		Updates(map[string]interface{}{
			"state":    "closed",
			"ended_at": &now,
		}).Error
}

func (dao *roomDAOImpl) Delete(roomID string) error {
	return dao.db.Delete(&model.Room{}, "room_id = ?", roomID).Error
}
func (dao *roomDAOImpl) CountRooms(userId string) (int64, error) {
	var count int64
	err := dao.db.Model(&model.Room{}).Where("host = ? and class_id IS NOT NULL and class_id !=''", userId).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
