package service

import repository "server/internal/handlers/repositories"

type ClassService struct {
	classRepo repository.ClassroomRepository
}

func NewClassService(classRepo repository.ClassroomRepository) *ClassService {
	return &ClassService{classRepo: classRepo}
}

func (s *ClassService) getClassroomsWithLimit(limit *int, ) {
}
