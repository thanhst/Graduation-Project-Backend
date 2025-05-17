package model

import (
	"time"
)

type User struct {
	UserId         string `gorm:"primaryKey;type:varchar(255);index"`
	FullName       string
	ProfilePicture string
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoCreateTime"`

	// Accounts      []Account      `gorm:"foreignKey:UserId;references:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	// Notifications []Notification `gorm:"foreignKey:UserId;references:UserId;"`
	// Student       Student        `gorm:"foreignKey:UserId;references:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	// Rooms         []Room         `gorm:"foreignKey:Host;references:UserId;"`
}
