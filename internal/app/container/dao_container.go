package container

import (
	accountdao "server/internal/handlers/dao/account"
	classroomdao "server/internal/handlers/dao/classroom"
	emotiondao "server/internal/handlers/dao/emotion"
	notificationdao "server/internal/handlers/dao/notification"
	roomdao "server/internal/handlers/dao/room"
	schedulerdao "server/internal/handlers/dao/scheduler"
	studentdao "server/internal/handlers/dao/student"
	studentclassdao "server/internal/handlers/dao/student_class"
	teacherdao "server/internal/handlers/dao/teacher"
	userdao "server/internal/handlers/dao/user"

	"gorm.io/gorm"
)

type DaoContainer struct {
	UserDAO         userdao.UserDAO
	AccountDAO      accountdao.AccountDAO
	RoomDAO         roomdao.RoomDAO
	StudentDAO      studentdao.StudentDAO
	StudentClassDAO studentclassdao.StudentClassDAO
	NotificationDAO notificationdao.NotificationDAO
	SchedulerDAO    schedulerdao.SchedulerDAO
	TeacherDAO      teacherdao.TeacherDAO
	EmotionDAO      emotiondao.EmotionDAO
	ClassroomDAO    classroomdao.ClassroomDAO
}

func NewDAOContainer(db *gorm.DB) *DaoContainer {
	return &DaoContainer{
		UserDAO:         userdao.NewUserDAO(db),
		AccountDAO:      accountdao.NewAccountDAO(db),
		RoomDAO:         roomdao.NewRoomDAO(db),
		StudentDAO:      studentdao.NewStudentDAO(db),
		StudentClassDAO: studentclassdao.NewStudentClassDAO(db),
		NotificationDAO: notificationdao.NewNotificationDAO(db),
		SchedulerDAO:    schedulerdao.NewSchedulerDAO(db),
		TeacherDAO:      teacherdao.NewTeacherDAO(db),
		EmotionDAO:      emotiondao.NewEmotionDAO(db),
		ClassroomDAO:    classroomdao.NewClassroomDAO(db),
	}
}
