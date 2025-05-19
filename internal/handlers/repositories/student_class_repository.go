package repository

import (
	studentclassdao "server/internal/handlers/dao/student_class"
	model "server/internal/models"
)

type StudentClassRepository interface {
	Get(userID, classID string) (*model.StudentClass, error)
	ListByUser(userID string) ([]model.StudentClass, error)
	ListByClass(classID string) ([]model.StudentClass, error)
	JoinClass(sc *model.StudentClass) error
	Update(sc *model.StudentClass) error
	Delete(userID, classID string) error
}

type studentClassRepository struct {
	dao studentclassdao.StudentClassDAO
}

func NewStudentClassRepository(dao studentclassdao.StudentClassDAO) StudentClassRepository {
	return &studentClassRepository{dao: dao}
}

func (r *studentClassRepository) Get(userID, classID string) (*model.StudentClass, error) {
	return r.dao.Get(userID, classID)
}

func (r *studentClassRepository) ListByUser(userID string) ([]model.StudentClass, error) {
	return r.dao.ListByUser(userID)
}

func (r *studentClassRepository) ListByClass(classID string) ([]model.StudentClass, error) {
	return r.dao.ListByClass(classID)
}

func (r *studentClassRepository) JoinClass(sc *model.StudentClass) error {
	return r.dao.JoinClass(sc)
}

func (r *studentClassRepository) Update(sc *model.StudentClass) error {
	return r.dao.Update(sc)
}

func (r *studentClassRepository) Delete(userID, classID string) error {
	return r.dao.Delete(userID, classID)
}
