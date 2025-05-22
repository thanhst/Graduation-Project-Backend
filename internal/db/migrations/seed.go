package migration

import (
	"context"
	"fmt"
	"math/rand"
	model "server/internal/models"
	CustomHash "server/internal/utils/hash"
	helper "server/internal/utils/helper"
	"time"

	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) error {
	ctx := context.Background()
	for i := 0; i < 100; i++ {
		user := model.User{
			UserId:         CustomHash.HashMD5(time.Now().String()),
			FullName:       helper.RandomName(),
			ProfilePicture: helper.RandomeImagesURL(),
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}
		if err := db.WithContext(ctx).Create(&user).Error; err != nil {
			return err
		}
		SeedAccounts(db, user)
	}
	return nil
}

func SeedAccounts(db *gorm.DB, user model.User) error {
	ctx := context.Background()
	role := helper.RandomRole()
	password, _ := CustomHash.HashPassword(user.FullName)
	account := model.Account{
		AccountId:   CustomHash.HashMD5(time.Now().String()),
		UserId:      user.UserId,
		Email:       helper.RandomEmail(),
		Password:    password,
		Role:        role,
		Status:      helper.RandomState(),
		LastLogin:   nil,
		LoginMethod: helper.RandomLoginMethod(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if err := db.WithContext(ctx).Create(&account).Error; err != nil {
		return err
	}
	if role == "teacher" {
		SeedTeacher(db, user.UserId)
	} else {
		SeedStudent(db, user.UserId)
	}
	return nil
}

func SeedTeacher(db *gorm.DB, userId string) error {
	ctx := context.Background()
	teacher := model.Teacher{
		UserId: userId,
	}
	if err := db.WithContext(ctx).Create(&teacher).Error; err != nil {
		return err
	}
	return nil
}

func SeedStudent(db *gorm.DB, userId string) error {
	ctx := context.Background()
	student := model.Student{
		UserId: userId,
	}
	if err := db.WithContext(ctx).Create(&student).Error; err != nil {
		return err
	}
	return nil
}

func SeedClassroom(db *gorm.DB) error {
	ctx := context.Background()
	var teachers []model.Teacher
	if err := db.WithContext(ctx).Preload("User").Find(&teachers).Error; err != nil {
		return err
	}
	if len(teachers) == 0 {
		return fmt.Errorf("no teacher found to assign classroom")
	}
	for i := 0; i < 100; i++ {
		randomTeacher := teachers[rand.Intn(len(teachers))]
		randomUser := randomTeacher.User.FullName
		classroom := model.Classroom{
			ClassId:     CustomHash.HashMD5(time.Now().String()),
			ClassName:   randomUser,
			UserCreated: randomTeacher.UserId,
			Description: helper.RandomDescription(),
			Url:         helper.RandomeImagesURL(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		if err := db.WithContext(ctx).Create(&classroom).Error; err != nil {
			return err
		}
		SeedStudentClassroom(db, classroom.ClassId)
	}
	return nil
}

func SeedStudentClassroom(db *gorm.DB, classId string) error {
	ctx := context.Background()
	var students []model.Student
	if err := db.WithContext(ctx).Preload("User").Find(&students).Error; err != nil {
		return err
	}
	var states = []string{
		"waiting",
		"joined",
	}
	randomStudent := students[rand.Intn(len(students))]
	randomUser := randomStudent.User.UserId
	classroom := model.StudentClass{
		ClassId:   classId,
		UserId:    randomUser,
		State:     states[rand.Intn(len(states))],
		CreatedAt: time.Now(),
		JoinedAt:  time.Now(),
	}
	if err := db.WithContext(ctx).Create(&classroom).Error; err != nil {
		return err
	}
	return nil
}
func SeedRoom(db *gorm.DB) error {
	ctx := context.Background()
	var classrooms []model.Classroom
	if err := db.WithContext(ctx).
		Preload("Teacher").
		Find(&classrooms).Error; err != nil {
		return err
	}
	var students []model.Student
	if err := db.WithContext(ctx).Preload("User").Find(&students).Error; err != nil {
		return err
	}
	for i := 0; i < 50; i++ {
		t := time.Now()
		randomClass := classrooms[rand.Intn(len(classrooms))]
		idClass := randomClass.ClassId
		roomMeeting := model.Room{
			RoomId:    CustomHash.HashMD5(time.Now().String()),
			ClassId:   &idClass,
			State:     "closed",
			Host:      randomClass.Teacher.UserId,
			CreatedAt: time.Now().Add(-24 * time.Hour),
			EndedAt:   &t,
		}
		if err := db.WithContext(ctx).Create(&roomMeeting).Error; err != nil {
			return err
		}
		SeedScheduler(db, roomMeeting, &randomClass.ClassId)
	}
	for i := 0; i < 50; i++ {
		t := time.Now()
		randomStudent := students[rand.Intn(len(students))]
		roomMeeting := model.Room{
			RoomId:    CustomHash.HashMD5(time.Now().String()),
			ClassId:   nil,
			State:     "closed",
			Host:      randomStudent.UserId,
			CreatedAt: time.Now().Add(-24 * time.Hour),
			EndedAt:   &t,
		}
		if err := db.WithContext(ctx).Create(&roomMeeting).Error; err != nil {
			return err
		}
		SeedScheduler(db, roomMeeting, nil)
	}
	return nil
}

func SeedScheduler(db *gorm.DB, randomRoom model.Room, classId *string) error {
	ctx := context.Background()
	scheduler := model.Scheduler{
		SchedulerId: CustomHash.HashMD5(time.Now().String()),
		RoomId:      randomRoom.RoomId,
		UserId:      randomRoom.Host,
		ClassId:     classId,
		StartTime:   time.Now().Add(+48 * time.Hour),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if err := db.WithContext(ctx).Create(&scheduler).Error; err != nil {
		return err
	}
	return nil
}

func SeedEmotion(db *gorm.DB) error {
	ctx := context.Background()
	var rooms []model.Room
	if err := db.WithContext(ctx).Find(&rooms).Error; err != nil {
		return err
	}
	var users []model.User
	if err := db.WithContext(ctx).Find(&users).Error; err != nil {
		return err
	}
	var emotions = []string{
		"Happy",
		"Fear",
		"Neutral",
		"Sad",
		"Surprise",
	}
	for i := 0; i < 100; i++ {
		emotion := model.Emotion{
			RoomId:    rooms[rand.Intn(len(rooms))].RoomId,
			UserId:    users[rand.Intn(len(users))].UserId,
			Emotion:   emotions[rand.Intn(len(emotions))],
			CreatedAt: time.Now(),
		}
		if err := db.WithContext(ctx).Create(&emotion).Error; err != nil {
			return err
		}
	}
	return nil
}

func SeedNotification(db *gorm.DB) error {
	ctx := context.Background()
	var users []model.User
	if err := db.WithContext(ctx).Find(&users).Error; err != nil {
		return err
	}
	var classes []model.Classroom
	if err := db.WithContext(ctx).Find(&classes).Error; err != nil {
		return err
	}
	for i := 0; i < 100; i++ {
		notification := model.Notification{
			NotificationId: CustomHash.HashMD5(time.Now().String()),
			UserId:         users[rand.Intn(len(users))].UserId,
			ClassId:        classes[rand.Intn(len(classes))].ClassId,
			Description:    helper.RandomDescription(),
			Type:           "info",
			CreatedAt:      time.Now(),
		}
		if err := db.WithContext(ctx).Create(&notification).Error; err != nil {
			return err
		}
	}
	return nil
}

func SeedAll(db *gorm.DB) bool {
	if err := SeedUsers(db); err != nil {
		return false
	}
	if err := SeedClassroom(db); err != nil {
		return false
	}
	if err := SeedRoom(db); err != nil {
		return false
	}
	if err := SeedEmotion(db); err != nil {
		return false
	}
	if err := SeedNotification(db); err != nil {
		return false
	}
	return true
}
