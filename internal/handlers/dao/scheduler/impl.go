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

func (dao *schedulerDAOImpl) GetSchedulerByUserAndDate(userId string, date string) ([]*model.Scheduler, error) {
	var schedulers []*model.Scheduler
	err := dao.db.
		Table("schedulers AS s").
		Joins("LEFT JOIN classrooms c ON s.class_id = c.class_id").
		Joins("LEFT JOIN student_classes sc ON sc.class_id = c.class_id").
		Preload("User").
		Preload("Classroom").
		Preload("Room").
		Where("DATE(s.start_time) = ?", date).
		Where("s.user_id = ? OR sc.user_id = ?", userId, userId).
		Select("DISTINCT s.*").
		Find(&schedulers).Error
	return schedulers, err
}
func (dao *schedulerDAOImpl) GetSchedulerByUser(userId string) ([]*model.Scheduler, error) {
	var schedulers []*model.Scheduler
	err := dao.db.
		Table("schedulers AS s").
		Joins("LEFT JOIN classrooms c ON s.class_id = c.class_id").
		Joins("LEFT JOIN student_classes sc ON sc.class_id = c.class_id").
		Preload("User").
		Preload("Classroom").
		Preload("Room").
		Where("s.user_id = ? OR sc.user_id = ?", userId, userId).
		Select("DISTINCT s.*").
		Find(&schedulers).Error
	return schedulers, err
}
func (dao *schedulerDAOImpl) View(sId string) (*model.Scheduler, error) {
	var scheduler *model.Scheduler
	err := dao.db.
		Preload("User").
		Preload("Classroom").
		Preload("Room").
		Where("scheduler_id = ?", sId).
		Find(&scheduler).Error
	return scheduler, err
}
func (dao *schedulerDAOImpl) GetCountSchedulerWithTime(classId string, date time.Time) (int64, error) {
	var count int64
	err := dao.db.
		Model(&model.Scheduler{}).
		Where("class_id = ? and start_time= ?", classId, date).
		Count(&count).Error
	return count, err
}
