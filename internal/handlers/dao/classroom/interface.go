package classroomdao

import model "server/internal/models"

type ClassroomDAO interface {
	GetByID(classID string) (*model.Classroom, error)
	GetAll() ([]*model.Classroom, error)
	GetByTeacherID(teacherID string, limit int, offset int) ([]*model.Classroom, error)
	Create(classroom *model.Classroom) error
	Update(classroom *model.Classroom) error
	Delete(classID string) error
	GetCountClassroomsByUser(userId string) (int64, error)
	GetClassroomsWithNewScheduler(userId string) ([]*model.Classroom, error)
}
