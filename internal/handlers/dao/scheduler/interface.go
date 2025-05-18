package schedulerdao

import model "server/internal/models"

type SchedulerDAO interface {
	GetByID(id string) (*model.Scheduler, error)
	GetByUserID(userID string, limit int, offset int) ([]model.Scheduler, error)
	GetByClassID(classID string) ([]model.Scheduler, error)
	Create(s *model.Scheduler) error
	Update(s *model.Scheduler) error
	Delete(id string) error
}
