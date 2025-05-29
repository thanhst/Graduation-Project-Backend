package repository

import (
	schedulerdao "server/internal/handlers/dao/scheduler"
	model "server/internal/models"
	"time"
)

type SchedulerRepository interface {
	GetByID(id string) (*model.Scheduler, error)
	GetByUserID(userID string, limit int, offset int) ([]*model.Scheduler, error)
	GetByClassID(classID string) ([]*model.Scheduler, error)
	Create(s *model.Scheduler) error
	Update(s *model.Scheduler) error
	Delete(schedulerId string) error
	GetSchedulerByUserAndDate(userId string, date time.Time) ([]*model.Scheduler, error)
}
type schedulerRepository struct {
	schedulerDAO schedulerdao.SchedulerDAO
}

func NewSchedulerRepository(dao schedulerdao.SchedulerDAO) SchedulerRepository {
	return &schedulerRepository{schedulerDAO: dao}
}

func (r *schedulerRepository) GetByID(id string) (*model.Scheduler, error) {
	return r.schedulerDAO.GetByID(id)
}

func (r *schedulerRepository) GetByUserID(userID string, limit int, offset int) ([]*model.Scheduler, error) {
	return r.schedulerDAO.GetByUserID(userID, limit, offset)
}

func (r *schedulerRepository) GetByClassID(classID string) ([]*model.Scheduler, error) {
	return r.schedulerDAO.GetByClassID(classID)
}

func (r *schedulerRepository) Create(s *model.Scheduler) error {
	return r.schedulerDAO.Create(s)
}

func (r *schedulerRepository) Update(s *model.Scheduler) error {
	return r.schedulerDAO.Update(s)
}

func (r *schedulerRepository) Delete(schedulerId string) error {
	return r.schedulerDAO.Delete(schedulerId)
}

func (r *schedulerRepository) GetSchedulerByUserAndDate(userId string, date time.Time) ([]*model.Scheduler, error) {
	return r.schedulerDAO.GetSchedulerByUserAndDate(userId, date)
}
