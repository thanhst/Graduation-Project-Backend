package teacherdao

import model "server/internal/models"

type TeacherDAO interface {
	GetByID(id uint) (*model.Teacher, error)
	GetByUserID(userId string) (*model.Teacher, error)
	GetAll() ([]model.Teacher, error)
	Create(teacher *model.Teacher) error
	Update(teacher *model.Teacher) error
	Delete(id uint) error
}
