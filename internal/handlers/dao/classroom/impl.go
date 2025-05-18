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

func (dao *classroomDAOImpl) GetAll() ([]model.Classroom, error) {
	var classrooms []model.Classroom
	err := dao.db.Preload("Teacher").Find(&classrooms).Error
	return classrooms, err
}

func (dao *classroomDAOImpl) GetByTeacherID(teacherID string, limit int, offset int) ([]model.Classroom, error) {
	var classrooms []model.Classroom
	err := dao.db.
		Where("user_created = ?", teacherID).
		Preload("Teacher").
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
