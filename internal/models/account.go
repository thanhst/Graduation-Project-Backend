package model

import (
	"time"
)

type Account struct {
	AccountId   string `gorm:"type:varchar(255);primaryKey;unique;not null;index"`
	UserId      string `gorm:"not null;type:varchar(255);index"`
	Email       string `gorm:"not null;"`
	Password    string `gorm:"not null"`
	Role        string `gorm:"type:enum('admin','teacher','student');not null"`
	Status      string `gorm:"type:enum('online','offline');default:'offline'"`
	LastLogin   *time.Time
	LoginMethod string    `gorm:"type:enum('google','github','email');default:'email'"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`

	User User `gorm:"foreignKey:UserId;references:UserId;"`
}
