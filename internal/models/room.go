package model

import (
	"time"
)

type Room struct {
	RoomId    string    `gorm:"type:varchar(255);primaryKey;not null"`
	ClassId   string    `gorm:"type:varchar(255);index"`
	State     string    `gorm:"type:enum('opening','closed')"`
	Host      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	EndedAt   *time.Time

	Classroom Classroom `gorm:"foreignKey:ClassId;references:ClassId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	User      User      `gorm:"foreignKey:Host;references:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
