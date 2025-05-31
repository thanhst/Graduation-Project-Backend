package service

import (
	repository "server/internal/handlers/repositories"
	model "server/internal/models"
)

type NotificationService struct {
	NotificationRepo repository.NotificationRepository
}

func NewNotificationService(NotificationRepo repository.NotificationRepository) *NotificationService {
	return &NotificationService{NotificationRepo: NotificationRepo}
}
func (NotificationService *NotificationService) Create(notification *model.Notification) error {
	return NotificationService.NotificationRepo.Create(notification)
}

//	func (NotificationService *NotificationService) Update(notification *model.Notification) error {
//		return NotificationService.NotificationRepo.Update(notification)
//	}
func (NotificationService *NotificationService) Delete(notificationId string) error {
	return NotificationService.NotificationRepo.Delete(notificationId)
}
func (NotificationService *NotificationService) GetByClasssrom(classId string) ([]*model.Notification, error) {
	return NotificationService.NotificationRepo.GetByClasssrom(classId)
}
