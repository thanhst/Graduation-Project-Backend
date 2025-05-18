package studentclassdao

import model "server/internal/models"
type StudentClassDAO interface {
	Get(userID, classID string) (*model.StudentClass, error)
	ListByUser(userID string) ([]model.StudentClass, error)
	ListByClass(classID string) ([]model.StudentClass, error)
	JoinClass(sc *model.StudentClass) error
	Update(sc *model.StudentClass) error
	Delete(userID, classID string) error
}