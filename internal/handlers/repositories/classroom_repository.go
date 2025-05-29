package repository

import (
	classroomdao "server/internal/handlers/dao/classroom"
	model "server/internal/models"
)

type ClassroomRepository interface {
	GetByID(classID string) (*model.Classroom, error)
	GetAll() ([]*model.Classroom, error)
	GetByTeacherID(teacherID string, limit int, offset int) ([]*model.Classroom, error)
	Create(classroom *model.Classroom) error
	Update(classroom *model.Classroom) error
	Delete(classID string) error
	GetCountClassroomsByUser(userId string) (int64, error)
	GetClassroomsWithNewScheduler(userId string) ([]*model.Classroom, error)
	GetTeacherFromClass(classId string) (*model.User, error)
}
type classroomRepository struct {
	classroomDAO classroomdao.ClassroomDAO
}

func NewClassroomRepository(dao classroomdao.ClassroomDAO) ClassroomRepository {
	return &classroomRepository{classroomDAO: dao}
}

func (r *classroomRepository) GetByID(classID string) (*model.Classroom, error) {
	return r.classroomDAO.GetByID(classID)
}

func (r *classroomRepository) GetAll() ([]*model.Classroom, error) {
	return r.classroomDAO.GetAll()
}

func (r *classroomRepository) GetByTeacherID(teacherID string, limit int, offset int) ([]*model.Classroom, error) {
	return r.classroomDAO.GetByTeacherID(teacherID, limit, offset)
}

func (r *classroomRepository) Create(classroom *model.Classroom) error {
	return r.classroomDAO.Create(classroom)
}

func (r *classroomRepository) Update(classroom *model.Classroom) error {
	return r.classroomDAO.Update(classroom)
}

func (r *classroomRepository) Delete(classID string) error {
	return r.classroomDAO.Delete(classID)
}

func (r *classroomRepository) GetCountClassroomsByUser(userId string) (int64, error) {
	return r.classroomDAO.GetCountClassroomsByUser(userId)
}
func (r *classroomRepository) GetClassroomsWithNewScheduler(userId string) ([]*model.Classroom, error) {
	return r.classroomDAO.GetClassroomsWithNewScheduler(userId)
}
func (r *classroomRepository) GetTeacherFromClass(classId string) (*model.User, error) {
	return r.classroomDAO.GetTeacherFromClass(classId)
}
