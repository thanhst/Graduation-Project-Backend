package studentclassdao

import model "server/internal/models"

type StudentClassDAO interface {
	Get(userID, classID string) (*model.StudentClass, error)
	ListByUser(userID string) ([]model.StudentClass, error)
	ListByClass(classID string) ([]model.StudentClass, error)
	JoinClass(sc *model.StudentClass) error
	Update(sc *model.StudentClass) error
	Delete(userID, classID string) error
	GetClassroomsWaitingByUser(userId string) ([]*model.Classroom, error)
	GetClassroomsJoinedByUser(userId string) ([]*model.Classroom, error)
	GetAllStudentWaitingWithClassroom(classId string) ([]*model.User, error)
	GetAllClassroomsByUser(userId string) ([]*model.Classroom, error)
	GetClassroomsByUser(userId string, limit int, offset int) ([]*model.Classroom, error)
	GetCountClassroomsByUser(userId string) (int64, error)
	GetClassroomsWithNewScheduler(userId string) ([]*model.Classroom, error)
}
