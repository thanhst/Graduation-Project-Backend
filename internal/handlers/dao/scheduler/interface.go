package schedulerdao

import (
	model "server/internal/models"
	"time"
)

type SchedulerDAO interface {
	GetByID(id string) (*model.Scheduler, error)
	GetByUserID(userID string, limit int, offset int) ([]*model.Scheduler, error)
	GetByClassID(classID string) ([]*model.Scheduler, error)
	Create(s *model.Scheduler) error
	Update(s *model.Scheduler) error
	Delete(id string) error
	GetSchedulerByUserAndDate(userId string, date string) ([]*model.Scheduler, error)
	GetSchedulerByUser(userId string) ([]*model.Scheduler, error)
	View(sId string) (*model.Scheduler, error)
	GetCountSchedulerWithTime(classId string, date time.Time) (int64, error)
}
