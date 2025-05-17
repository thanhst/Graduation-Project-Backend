package migration

import (
	model "server/internal/models"

	"gorm.io/gorm"
)

func MigrateWithGORM(db *gorm.DB) error {
	if err := db.AutoMigrate(&model.User{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&model.Teacher{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&model.Classroom{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&model.Student{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&model.Account{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&model.Notification{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&model.Room{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&model.Scheduler{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&model.StudentClass{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&model.Emotion{}); err != nil {
		return err
	}
	return nil
}
