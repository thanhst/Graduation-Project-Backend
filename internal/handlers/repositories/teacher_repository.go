package repository

import (
	teacherdao "server/internal/handlers/dao/teacher"
	model "server/internal/models"
)

type TeacherRepository interface {
	GetByID(id uint) (*model.Teacher, error)
	GetByUserID(userId string) (*model.Teacher, error)
	GetAll() ([]model.Teacher, error)
	Create(teacher *model.Teacher) error
	Update(teacher *model.Teacher) error
	Delete(id uint) error
}
type teacherRepository struct {
	teacherDAO teacherdao.TeacherDAO
}

func NewTeacherRepository(dao teacherdao.TeacherDAO) TeacherRepository {
	return &teacherRepository{teacherDAO: dao}
}

func (r *teacherRepository) GetByID(id uint) (*model.Teacher, error) {
	return r.teacherDAO.GetByID(id)
}

func (r *teacherRepository) GetByUserID(userId string) (*model.Teacher, error) {
	return r.teacherDAO.GetByUserID(userId)
}

func (r *teacherRepository) GetAll() ([]model.Teacher, error) {
	return r.teacherDAO.GetAll()
}

func (r *teacherRepository) Create(teacher *model.Teacher) error {
	return r.teacherDAO.Create(teacher)
}

func (r *teacherRepository) Update(teacher *model.Teacher) error {
	return r.teacherDAO.Update(teacher)
}

func (r *teacherRepository) Delete(id uint) error {
	return r.teacherDAO.Delete(id)
}
