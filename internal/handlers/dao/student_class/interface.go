package studentclassdao

import model "server/internal/models"

type StudentClassDAO interface {
	Get(userID, classID string) (*model.StudentClass, error)
	ListByUser(userID string) ([]model.StudentClass, error)
	ListByClass(classID string) ([]model.StudentClass, error)
	JoinClass(sc *model.StudentClass) error
	Update(sc *model.StudentClass) error
	Delete(sc *model.StudentClass) error
	GetClassroomsWaitingByUser(userId string) ([]*model.Classroom, error)
	GetClassroomsJoinedByUser(userId string) ([]*model.Classroom, error)
	GetAllStudentWaitingWithClassroom(classId string) ([]*model.User, error)
	GetAllClassroomsByUser(userId string) ([]*model.Classroom, error)
	GetClassroomsByUser(userId string, limit int, offset int) ([]*model.Classroom, error)
	GetCountClassroomsByUser(userId string) (int64, error)
	GetClassroomsWithNewScheduler(userId string) ([]*model.Classroom, error)
	GetCountUsersByClassroom(classId string) (int64, int64, error)
	GetUserWaitingWithClassrooms(classId string, limit int, offset int) ([]*model.User, error)
	GetUserJoinedWithClassrooms(classId string, limit int, offset int) ([]*model.User, error)
	GetInfo(classId string, userId string) (*model.StudentClass, error)
}
