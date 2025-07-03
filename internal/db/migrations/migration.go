package migration

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"
)

func MigrateWithMigration(gdb *gorm.DB) error {
	sqlDB, err := gdb.DB()
	if err != nil {
		return err
	}

	driver, err := mysql.WithInstance(sqlDB, &mysql.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/db/migrations/sql",
		"mysql",
		driver,
	)
	if err != nil {
		return err
	}
	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		log.Println("Warning during down:", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	log.Println("Migration success!")
	return nil
}
