package model

import (
	"time"
)

type Room struct {
	RoomId    string     `gorm:"type:varchar(255);primaryKey;not null" json:"roomID"`
	ClassId   *string    `gorm:"type:varchar(255);" json:"classID,omitempty"`
	State     string     `gorm:"type:enum('opening','closed');default:closed" json:"state"`
	Host      string     `gorm:"not null" json:"host"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"createdAt"`
	EndedAt   *time.Time `json:"endAt,omitempty"`

	Classroom Classroom `gorm:"foreignKey:ClassId;references:ClassId;"`
	User      User      `gorm:"foreignKey:Host;references:UserId;"`
}
