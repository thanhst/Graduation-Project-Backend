package studentdao

import (
	model "server/internal/models"

	"gorm.io/gorm"
)

type studentDAOImpl struct {
	db *gorm.DB
}

func NewStudentDAO(db *gorm.DB) StudentDAO {
	return &studentDAOImpl{db: db}
}

func (dao *studentDAOImpl) GetByUserID(userID string) (*model.Student, error) {
	var student model.Student
	err := dao.db.
		Preload("User").
		Where("user_id = ?", userID).
		First(&student).Error
	return &student, err
}

func (dao *studentDAOImpl) Create(student *model.Student) error {
	return dao.db.Create(student).Error
}

func (dao *studentDAOImpl) Update(student *model.Student) error {
	return dao.db.Save(student).Error
}

func (dao *studentDAOImpl) Delete(userID string) error {
	return dao.db.Where("user_id = ?", userID).Delete(&model.Student{}).Error
}
