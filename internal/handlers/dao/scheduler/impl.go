package schedulerdao

import (
	model "server/internal/models"
	"time"

	"gorm.io/gorm"
)

type schedulerDAOImpl struct {
	db *gorm.DB
}

func NewSchedulerDAO(db *gorm.DB) SchedulerDAO {
	return &schedulerDAOImpl{db: db}
}

func (dao *schedulerDAOImpl) GetByID(id string) (*model.Scheduler, error) {
	var s model.Scheduler
	err := dao.db.
		Preload("Classroom").
		Preload("Room").
		Preload("User").
		First(&s, "scheduler_id = ?", id).Error
	return &s, err
}

func (dao *schedulerDAOImpl) GetByUserID(userID string, limit int, offset int) ([]*model.Scheduler, error) {
	var schedulers []*model.Scheduler
	err := dao.db.
		Preload("Room").
		Preload("Classroom").
		Where("user_id = ?", userID).
		Order("start_time ASC").
		Limit(limit).
		Offset(offset).
		Find(&schedulers).Error
	return schedulers, err
}

func (dao *schedulerDAOImpl) GetByClassID(classID string) ([]*model.Scheduler, error) {
	var schedulers []*model.Scheduler
	err := dao.db.
		Where("class_id = ?", classID).
		Find(&schedulers).Error
	return schedulers, err
}

func (dao *schedulerDAOImpl) Create(s *model.Scheduler) error {
	return dao.db.Create(s).Error
}

func (dao *schedulerDAOImpl) Update(s *model.Scheduler) error {
	return dao.db.Save(s).Error
}

func (dao *schedulerDAOImpl) Delete(id string) error {
	return dao.db.Delete(&model.Scheduler{}, "scheduler_id = ?", id).Error
}

func (dao *schedulerDAOImpl) GetSchedulerByUserAndDate(userId string, date time.Time) ([]*model.Scheduler, error) {
	var schedulers []*model.Scheduler
	err := dao.db.
		Where("user_id = ? and start_time =? ", userId, date).
		Find(&schedulers).Error
	return schedulers, err
}
