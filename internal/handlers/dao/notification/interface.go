package notificationdao

import model "server/internal/models"

type NotificationDAO interface {
	GetByID(id string) (*model.Notification, error)
	GetByUser(userId string, limit int, offset int) ([]model.Notification, error)
	Create(notification *model.Notification) error
	Delete(id string) error
	DeleteAllOfUser(userId string) error
}
