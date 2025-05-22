package container

import service "server/internal/handlers/services"

type ServiceContainer struct {
	UserService    *service.UserService
	AccountService *service.AccountService
	// RoomService         service.RoomService
	// StudentService      service.StudentService
	// StudentClassService service.StudentClassService
	// NotificationService service.NotificationService
	// SchedulerService    service.SchedulerService
	// TeacherService      service.TeacherService
	// EmotionService      service.EmotionService
	// ClassroomService    service.ClassroomService
}

func NewServiceContainer(repoContainer *RepoContainer) *ServiceContainer {
	return &ServiceContainer{
		UserService:    service.NewUserService(repoContainer.UserRepo),
		AccountService: service.NewAccountService(repoContainer.AccountRepo),
		// RoomService:         service.NewRoomService(repoContainer.RoomRepo),
		// StudentService:      service.NewStudentService(repoContainer.StudentRepo),
		// StudentClassService: service.NewStudentClassService(repoContainer.StudentClassRepo),
		// NotificationService: service.NewNotificationService(repoContainer.NotificationRepo),
		// SchedulerService:    service.NewSchedulerService(repoContainer.SchedulerRepo),
		// TeacherService:      service.NewTeacherService(repoContainer.TeacherRepo),
		// EmotionService:      service.NewEmotionService(repoContainer.EmotionRepo),
		// ClassroomService:    service.NewClassroomService(repoContainer.ClassroomRepo),
	}
}
