package studentdao

import model "server/internal/models"

type StudentDAO interface {
	GetByUserID(userID string) (*model.Student, error)
	Create(student *model.Student) error
	Update(student *model.Student) error
	Delete(userID string) error
}
