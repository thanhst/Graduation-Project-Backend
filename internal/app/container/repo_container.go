package container

import repository "server/internal/handlers/repositories"

type RepoContainer struct {
	UserRepo         repository.UserRepository
	AccountRepo      repository.AccountRepository
	RoomRepo         repository.RoomRepository
	StudentRepo      repository.StudentRepository
	StudentClassRepo repository.StudentClassRepository
	NotificationRepo repository.NotificationRepository
	SchedulerRepo    repository.SchedulerRepository
	TeacherRepo      repository.TeacherRepository
	EmotionRepo      repository.EmotionRepository
	ClassroomRepo    repository.ClassroomRepository
}

func NewRepoContainer(daoContainer *DaoContainer) *RepoContainer {
	return &RepoContainer{
		UserRepo:         repository.NewUserRepository(daoContainer.UserDAO),
		AccountRepo:      repository.NewAccountRepository(daoContainer.AccountDAO),
		RoomRepo:         repository.NewRoomRepository(daoContainer.RoomDAO),
		StudentRepo:      repository.NewStudentRepository(daoContainer.StudentDAO),
		StudentClassRepo: repository.NewStudentClassRepository(daoContainer.StudentClassDAO),
		NotificationRepo: repository.NewNotificationRepository(daoContainer.NotificationDAO),
		SchedulerRepo:    repository.NewSchedulerRepository(daoContainer.SchedulerDAO),
		TeacherRepo:      repository.NewTeacherRepository(daoContainer.TeacherDAO),
		EmotionRepo:      repository.NewEmotionRepository(daoContainer.EmotionDAO),
		ClassroomRepo:    repository.NewClassroomRepository(daoContainer.ClassroomDAO),
	}
}
