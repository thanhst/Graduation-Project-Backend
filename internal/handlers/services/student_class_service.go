package service

import (
	repository "server/internal/handlers/repositories"
	model "server/internal/models"
)

type StudentClassService struct {
	studentClassRepo repository.StudentClassRepository
}

func NewStudentClassService(studentClassRepo repository.StudentClassRepository) *StudentClassService {
	return &StudentClassService{studentClassRepo: studentClassRepo}
}

func (stcls *StudentClassService) GetClassroomsByUser(userId string, limit int, offset int) ([]*model.Classroom, error) {
	return stcls.studentClassRepo.GetClassroomsByUser(userId, limit, offset)
}

func (stcls *StudentClassService) GetCountClassroomsByUser(userId string) (int64, error) {
	return stcls.studentClassRepo.GetCountClassroomsByUser(userId)
}
func (stcls *StudentClassService) GetClassroomsWithNewScheduler(userId string) ([]*model.Classroom, error) {
	return stcls.studentClassRepo.GetClassroomsWithNewScheduler(userId)
}

// func (stcls *StudentClassService) GetAllClassroomsByUser(userId string) (int64, error) {
// 	return stcls.studentClassRepo.GetAllClassroomsByUser(userId)
// }
