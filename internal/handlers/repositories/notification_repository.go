package repository

import (
	notificationdao "server/internal/handlers/dao/notification"
	model "server/internal/models"
)

type NotificationRepository interface {
	GetByID(id string) (*model.Notification, error)
	GetByUser(userID string, limit int, offset int) ([]model.Notification, error)
	Create(noti *model.Notification) error
	Delete(id string) error
	DeleteAllOfUser(userID string) error
	GetByClasssrom(classId string) ([]*model.Notification, error)
	GetByUserClassrooms(userId string) ([]*model.Notification, error)
	GetLatestByUserClassrooms(userId string) (*model.Notification, error)
}
type notificationRepository struct {
	dao notificationdao.NotificationDAO
}

func NewNotificationRepository(dao notificationdao.NotificationDAO) NotificationRepository {
	return &notificationRepository{dao: dao}
}

func (r *notificationRepository) GetByID(id string) (*model.Notification, error) {
	return r.dao.GetByID(id)
}

func (r *notificationRepository) GetByUser(userID string, limit int, offset int) ([]model.Notification, error) {
	return r.dao.GetByUser(userID, limit, offset)
}

func (r *notificationRepository) Create(noti *model.Notification) error {
	return r.dao.Create(noti)
}

func (r *notificationRepository) Delete(id string) error {
	return r.dao.Delete(id)
}

func (r *notificationRepository) DeleteAllOfUser(userID string) error {
	return r.dao.DeleteAllOfUser(userID)
}
func (r *notificationRepository) GetByClasssrom(userId string) ([]*model.Notification, error) {
	return r.dao.GetByClasssrom(userId)
}

func (r *notificationRepository) GetByUserClassrooms(userId string) ([]*model.Notification, error) {
	return r.dao.GetByUserClassrooms(userId)
}
func (r *notificationRepository) GetLatestByUserClassrooms(userId string) (*model.Notification, error) {
	return r.dao.GetLatestByUserClassrooms(userId)
}
