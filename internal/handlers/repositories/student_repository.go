package repository

import (
	studentdao "server/internal/handlers/dao/student"
	model "server/internal/models"
)

type StudentRepository interface {
	GetByUserID(userID string) (*model.Student, error)
	Create(student *model.Student) error
	Update(student *model.Student) error
	Delete(userID string) error
}

type studentRepository struct {
	studentDAO studentdao.StudentDAO
}

func NewStudentRepository(dao studentdao.StudentDAO) StudentRepository {
	return &studentRepository{studentDAO: dao}
}

func (r *studentRepository) GetByUserID(userID string) (*model.Student, error) {
	return r.studentDAO.GetByUserID(userID)
}

func (r *studentRepository) Create(student *model.Student) error {
	return r.studentDAO.Create(student)
}

func (r *studentRepository) Update(student *model.Student) error {
	return r.studentDAO.Update(student)
}

func (r *studentRepository) Delete(userID string) error {
	return r.studentDAO.Delete(userID)
}
