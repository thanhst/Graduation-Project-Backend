package database

import (
	"database/sql"
	"fmt"
	"os"
	migration "server/internal/db/migrations"
	"server/internal/utils/dotenv"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GDB *gorm.DB

func createDatabaseIfNotExist(user, pass, host, port, dbname string) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", user, pass, host, port)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;", dbname)
	_, err = db.Exec(query)
	return err
}

func ConnectDB() error {
	if err := godotenv.Load("cmd/config/.env"); err != nil {
		fmt.Println("Warning: .env file not found!")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")

	err := createDatabaseIfNotExist(user, pass, host, port, dbname)
	if err != nil {
		return fmt.Errorf("failed to create db: %w", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, dbname,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}

	GDB = db

	schemaVer := dotenv.GetDotEnv("SCHEMA_VER")
	schemaVerInt, err := strconv.Atoi(schemaVer)
	if err != nil {
		return fmt.Errorf("error to convert schema version: %w", err)
	}
	if schemaVerInt == 0 {
		err = migration.MigrateWithMigration(GDB)
		if err != nil {
			return fmt.Errorf("error to migration: %w", err)
		}

		if dotenv.GetDotEnv("APP_ENV") == "develop" {
			if checkFlag := migration.SeedAll(GDB); !checkFlag {
				return fmt.Errorf("error to seed data")
			}
		}
		err = dotenv.SetDotEnv("SCHEMA_VER", "1")
		if err != nil {
			return fmt.Errorf("error updating APP_ENV: %v", err)
		}
		fmt.Println("Seed success to database!")
	}

	fmt.Println("Connect success to database")
	return nil
}
