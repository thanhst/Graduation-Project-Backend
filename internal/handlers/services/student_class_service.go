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

func (stcls *StudentClassService) JoinClass(stdClass *model.StudentClass) error {
	return stcls.studentClassRepo.JoinClass(stdClass)
}
func (stcls *StudentClassService) GetUserJoinedWithClassrooms(classId string, limit int, offset int) ([]*model.User, error) {
	return stcls.studentClassRepo.GetUserJoinedWithClassrooms(classId, limit, offset)
}
func (stcls *StudentClassService) GetUserWaitingWithClassrooms(classId string, limit int, offset int) ([]*model.User, error) {
	return stcls.studentClassRepo.GetUserWaitingWithClassrooms(classId, limit, offset)
}
func (stcls *StudentClassService) GetCountUsersByClassroom(classId string) (int64, int64, error) {
	return stcls.studentClassRepo.GetCountUsersByClassroom(classId)
}
func (stcls *StudentClassService) Update(std *model.StudentClass) error {
	return stcls.studentClassRepo.Update(std)
}
func (stcls *StudentClassService) Delete(std *model.StudentClass) error {
	return stcls.studentClassRepo.Delete(std)
}
func (stcls *StudentClassService) GetInfo(classId string, userId string) (*model.StudentClass, error) {
	return stcls.studentClassRepo.GetInfo(classId, userId)
}
