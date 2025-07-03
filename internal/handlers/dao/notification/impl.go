package notificationdao

import (
	model "server/internal/models"

	"gorm.io/gorm"
)

type notificationDAOImpl struct {
	db *gorm.DB
}

func NewNotificationDAO(db *gorm.DB) NotificationDAO {
	return &notificationDAOImpl{db: db}
}

func (dao *notificationDAOImpl) GetByID(id string) (*model.Notification, error) {
	var noti model.Notification
	err := dao.db.
		Preload("User").
		Preload("Classroom").
		First(&noti, "notification_id = ?", id).Error
	return &noti, err
}

func (dao *notificationDAOImpl) GetByUser(userId string, limit int, offset int) ([]model.Notification, error) {
	var notis []model.Notification
	err := dao.db.
		Where("user_id = ?", userId).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Preload("Classroom").
		Find(&notis).Error
	return notis, err
}

func (dao *notificationDAOImpl) Create(noti *model.Notification) error {
	return dao.db.Create(noti).Error
}

func (dao *notificationDAOImpl) Delete(id string) error {
	return dao.db.Delete(&model.Notification{}, "notification_id = ?", id).Error
}

func (dao *notificationDAOImpl) DeleteAllOfUser(userId string) error {
	return dao.db.Where("user_id = ?", userId).Delete(&model.Notification{}).Error
}

func (dao *notificationDAOImpl) GetByClasssrom(classID string) ([]*model.Notification, error) {
	var data []*model.Notification
	if err := dao.db.Where("class_id = ?", classID).
		Order("created_at DESC").
		Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
func (dao *notificationDAOImpl) GetByUserClassrooms(userID string) ([]*model.Notification, error) {
	var data []*model.Notification

	err := dao.db.
		Table("notifications AS n").
		Joins("LEFT JOIN classrooms c ON n.class_id = c.class_id").
		Joins("LEFT JOIN student_classes sc ON sc.class_id = c.class_id AND sc.user_id = ?", userID).
		Where("sc.user_id IS NOT NULL OR n.user_id = ?", userID).
		Order("n.created_at DESC").
		Preload("Scheduler").
		Preload("Classroom").
		Find(&data).Error
	return data, err
}
func (dao *notificationDAOImpl) GetLatestByUserClassrooms(userID string) (*model.Notification, error) {
	var notif model.Notification

	err := dao.db.
		Table("notifications AS n").
		Joins("LEFT JOIN classrooms c ON n.class_id = c.class_id").
		Joins("LEFT JOIN student_classes sc ON sc.class_id = c.class_id AND sc.user_id = ?", userID).
		Where("sc.user_id IS NOT NULL OR n.user_id = ?", userID).
		Order("n.created_at DESC").
		Limit(1).
		Preload("Classroom").
		Preload("Scheduler").
		Find(&notif).Error

	if err != nil {
		return nil, err
	}
	return &notif, nil
}
