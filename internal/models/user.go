package model

import (
	"time"
)

type User struct {
	UserId         string    `gorm:"primaryKey;type:varchar(255);index" json:"userId"`
	FullName       string    `json:"fullName"`
	ProfilePicture string    `json:"profilePicture"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt      time.Time `gorm:"autoCreateTime" json:"updatedAt"`

	Accounts      []Account      `gorm:"foreignKey:UserId;references:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Notifications []Notification `gorm:"foreignKey:UserId;references:UserId;"`
	Rooms         []Room         `gorm:"foreignKey:Host;references:UserId;"`
}
