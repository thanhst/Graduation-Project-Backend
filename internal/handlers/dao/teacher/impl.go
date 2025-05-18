package teacherdao

import (
	model "server/internal/models"

	"gorm.io/gorm"
)

type teacherDAOImpl struct {
	db *gorm.DB
}

func NewTeacherDAO(db *gorm.DB) TeacherDAO {
	return &teacherDAOImpl{db: db}
}

func (dao *teacherDAOImpl) GetByID(id uint) (*model.Teacher, error) {
	var teacher model.Teacher
	err := dao.db.
		Preload("User").
		Preload("Classrooms").
		First(&teacher, id).Error
	return &teacher, err
}

func (dao *teacherDAOImpl) GetByUserID(userId string) (*model.Teacher, error) {
	var teacher model.Teacher
	err := dao.db.
		Preload("User").
		Preload("Classrooms").
		Where("user_id = ?", userId).
		First(&teacher).Error
	return &teacher, err
}

func (dao *teacherDAOImpl) GetAll() ([]model.Teacher, error) {
	var teachers []model.Teacher
	err := dao.db.
		Preload("User").
		Preload("Classrooms").
		Find(&teachers).Error
	return teachers, err
}

func (dao *teacherDAOImpl) Create(teacher *model.Teacher) error {
	return dao.db.Create(teacher).Error
}

func (dao *teacherDAOImpl) Update(teacher *model.Teacher) error {
	return dao.db.Save(teacher).Error
}

func (dao *teacherDAOImpl) Delete(id uint) error {
	return dao.db.Delete(&model.Teacher{}, id).Error
}
