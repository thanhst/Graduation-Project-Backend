package service

import (
	repository "server/internal/handlers/repositories"
	model "server/internal/models"
)

type ClassService struct {
	classRepo repository.ClassroomRepository
}

func NewClassService(classRepo repository.ClassroomRepository) *ClassService {
	return &ClassService{classRepo: classRepo}
}

func (s *ClassService) GetClassroomsWithLimit(limit *int) {
}

func (s *ClassService) GetClassroomsByUser(userId string, limit int, offset int) ([]*model.Classroom, error) {
	return s.classRepo.GetByTeacherID(userId, limit, offset)
}
func (s *ClassService) GetCountClassroomsByUser(userId string) (int64, error) {
	return s.classRepo.GetCountClassroomsByUser(userId)
}
func (s *ClassService) GetClassroomsWithNewScheduler(userId string) ([]*model.Classroom, error) {
	return s.classRepo.GetClassroomsWithNewScheduler(userId)
}
func (s *ClassService) GetClassroomById(classId string) (*model.Classroom, error) {
	return s.classRepo.GetClassroomById(classId)
}
func (s *ClassService) Create(classroom *model.Classroom) error {
	return s.classRepo.Create(classroom)
}
func (s *ClassService) Update(classroom *model.Classroom) error {
	return s.classRepo.Update(classroom)
}
func (s *ClassService) Delete(classId string) error {
	return s.classRepo.Delete(classId)
}
