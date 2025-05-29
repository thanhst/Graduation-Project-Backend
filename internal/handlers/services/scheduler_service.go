package service

import (
	repository "server/internal/handlers/repositories"
	model "server/internal/models"
	"time"
)

type SchedulerService struct {
	SchedulerRepo repository.SchedulerRepository
}

func NewSchedulerService(SchedulerRepo repository.SchedulerRepository) *SchedulerService {
	return &SchedulerService{SchedulerRepo: SchedulerRepo}
}

func (scheduler *SchedulerService) Create(schedule *model.Scheduler) error {
	return scheduler.SchedulerRepo.Create(schedule)
}
func (scheduler *SchedulerService) Update(schedule *model.Scheduler) error {
	return scheduler.SchedulerRepo.Update(schedule)
}
func (scheduler *SchedulerService) Delete(scheduleId string) error {
	return scheduler.SchedulerRepo.Delete(scheduleId)
}

func (scheduler *SchedulerService) GetSchedulerByUserAndDate(userId string, date time.Time) ([]*model.Scheduler, error) {
	return scheduler.SchedulerRepo.GetSchedulerByUserAndDate(userId, date)
}

