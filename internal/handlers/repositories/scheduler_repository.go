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
	GetSchedulerByUserAndDate(userId string, date string) ([]*model.Scheduler, error)
	GetSchedulerByUser(userId string) ([]*model.Scheduler, error)
	View(sId string) (*model.Scheduler, error)
	GetCountSchedulerWithTime(classId string, date time.Time) (int64, error)
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

func (r *schedulerRepository) GetSchedulerByUserAndDate(userId string, date string) ([]*model.Scheduler, error) {
	return r.schedulerDAO.GetSchedulerByUserAndDate(userId, date)
}

func (r *schedulerRepository) GetSchedulerByUser(userId string) ([]*model.Scheduler, error) {
	return r.schedulerDAO.GetSchedulerByUser(userId)
}
func (r *schedulerRepository) View(sId string) (*model.Scheduler, error) {
	return r.schedulerDAO.View(sId)
}
func (r *schedulerRepository) GetCountSchedulerWithTime(classId string, date time.Time) (int64, error) {
	return r.schedulerDAO.GetCountSchedulerWithTime(classId, date)
}
