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
	Delete(sc *model.StudentClass) error
	GetClassroomsByUser(userId string, limit int, offset int) ([]*model.Classroom, error)
	GetCountClassroomsByUser(userId string) (int64, error)
	GetClassroomsWithNewScheduler(userId string) ([]*model.Classroom, error)
	GetUserJoinedWithClassrooms(classId string, limit int, offset int) ([]*model.User, error)
	GetUserWaitingWithClassrooms(classId string, limit int, offset int) ([]*model.User, error)
	GetCountUsersByClassroom(classId string) (int64, int64, error)
	GetInfo(classId string, userId string) (*model.StudentClass, error)
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

func (r *studentClassRepository) Delete(sc *model.StudentClass) error {
	return r.dao.Delete(sc)
}

func (r *studentClassRepository) GetClassroomsByUser(userId string, limit int, offset int) ([]*model.Classroom, error) {
	return r.dao.GetClassroomsByUser(userId, limit, offset)
}
func (r *studentClassRepository) GetCountClassroomsByUser(userId string) (int64, error) {
	return r.dao.GetCountClassroomsByUser(userId)
}
func (r *studentClassRepository) GetClassroomsWithNewScheduler(userId string) ([]*model.Classroom, error) {
	return r.dao.GetClassroomsWithNewScheduler(userId)
}
func (r *studentClassRepository) GetUserJoinedWithClassrooms(classId string, limit int, offset int) ([]*model.User, error) {
	return r.dao.GetUserJoinedWithClassrooms(classId, limit, offset)
}
func (r *studentClassRepository) GetUserWaitingWithClassrooms(classId string, limit int, offset int) ([]*model.User, error) {
	return r.dao.GetUserWaitingWithClassrooms(classId, limit, offset)
}
func (r *studentClassRepository) GetCountUsersByClassroom(classId string) (int64, int64, error) {
	return r.dao.GetCountUsersByClassroom(classId)
}
func (r *studentClassRepository) GetInfo(classId string, userId string) (*model.StudentClass, error) {
	return r.dao.GetInfo(classId, userId)
}
