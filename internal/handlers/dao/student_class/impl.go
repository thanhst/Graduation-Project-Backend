package studentclassdao

import (
	model "server/internal/models"

	"gorm.io/gorm"
)

type studentClassDAOImpl struct {
	db *gorm.DB
}

func NewStudentClassDAO(db *gorm.DB) StudentClassDAO {
	return &studentClassDAOImpl{db: db}
}

func (dao *studentClassDAOImpl) Get(userID, classID string) (*model.StudentClass, error) {
	var sc model.StudentClass
	err := dao.db.
		Preload("User").
		Preload("Classroom").
		Where("user_id = ? AND class_id = ?", userID, classID).
		First(&sc).Error
	return &sc, err
}

func (dao *studentClassDAOImpl) ListByUser(userID string) ([]model.StudentClass, error) {
	var scs []model.StudentClass
	err := dao.db.
		Preload("Classroom").
		Where("user_id = ?", userID).
		Find(&scs).Error
	return scs, err
}

func (dao *studentClassDAOImpl) ListByClass(classID string) ([]model.StudentClass, error) {
	var scs []model.StudentClass
	err := dao.db.
		Preload("User").
		Where("class_id = ?", classID).
		Find(&scs).Error
	return scs, err
}

func (dao *studentClassDAOImpl) JoinClass(sc *model.StudentClass) error {
	return dao.db.Create(sc).Error
}

func (dao *studentClassDAOImpl) Update(sc *model.StudentClass) error {
	return dao.db.Save(sc).Error
}

func (dao *studentClassDAOImpl) Delete(userID, classID string) error {
	return dao.db.Where("user_id = ? AND class_id = ?", userID, classID).
		Delete(&model.StudentClass{}).Error
}
