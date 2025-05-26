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
