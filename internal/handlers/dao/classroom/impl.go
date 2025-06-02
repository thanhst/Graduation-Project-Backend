package classroomdao

import (
	model "server/internal/models"

	"gorm.io/gorm"
)

type classroomDAOImpl struct {
	db *gorm.DB
}

func NewClassroomDAO(db *gorm.DB) ClassroomDAO {
	return &classroomDAOImpl{db: db}
}

func (dao *classroomDAOImpl) GetByID(classID string) (*model.Classroom, error) {
	var classroom model.Classroom
	err := dao.db.Preload("Teacher").First(&classroom, "class_id = ?", classID).Error
	return &classroom, err
}

func (dao *classroomDAOImpl) GetAll() ([]*model.Classroom, error) {
	var classrooms []*model.Classroom
	err := dao.db.Preload("Teacher").Find(&classrooms).Error
	return classrooms, err
}

func (dao *classroomDAOImpl) GetByTeacherID(teacherID string, limit int, offset int) ([]*model.Classroom, error) {
	var classrooms []*model.Classroom
	err := dao.db.
		Where("user_created = ?", teacherID).
		Preload("User").
		Preload("StudentClasses", "state=?", "joined").Preload("StudentClasses.User").
		Preload("Notifications").
		Preload("Schedulers").
		Limit(limit).
		Offset(offset).
		Find(&classrooms).Error
	return classrooms, err
}

func (dao *classroomDAOImpl) Create(classroom *model.Classroom) error {
	return dao.db.Create(classroom).Error
}

func (dao *classroomDAOImpl) Update(classroom *model.Classroom) error {
	return dao.db.Save(classroom).Error
}

func (dao *classroomDAOImpl) Delete(classID string) error {
	return dao.db.Delete(&model.Classroom{}, "class_id = ?", classID).Error
}

func (dao *classroomDAOImpl) GetCountClassroomsByUser(userId string) (int64, error) {
	var count int64
	err := dao.db.Model(&model.Classroom{}).Where("user_created = ?", userId).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (dao *classroomDAOImpl) GetClassroomsWithNewScheduler(userId string) ([]*model.Classroom, error) {
	var classrooms []*model.Classroom

	subQuery := dao.db.Model(&model.Scheduler{}).
		Select("class_id, MIN(start_time) AS nearest_start").
		Where("DATE(start_time) = CURDATE()").
		Group("class_id")

	err := dao.db.Model(&model.Classroom{}).
		Joins("LEFT JOIN (?) AS smin ON smin.class_id = classrooms.class_id", subQuery).
		Where("classrooms.user_created = ? AND smin.nearest_start IS NOT NULL", userId).
		Preload("Schedulers").
		Preload("StudentClasses.User").
		Preload("User").
		Find(&classrooms).Error

	if err != nil {
		return nil, err
	}
	return classrooms, nil
}

func (dao *classroomDAOImpl) GetClassroomById(classId string) (*model.Classroom, error) {
	var classroom model.Classroom
	err := dao.db.Where("class_id = ?", classId).
		Preload("User").
		Preload("Notifications").
		Preload("Schedulers").
		Preload("StudentClasses").
		Find(&classroom).Error
	if err != nil {
		return nil, err
	}
	return &classroom, nil
}
