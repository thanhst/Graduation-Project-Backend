package model

import (
	"time"
)

type Scheduler struct {
	SchedulerId string    `gorm:"type:varchar(255);primaryKey;not null;index"`
	RoomId      string    `gorm:"type:varchar(255);not null;index"`
	UserId      string    `gorm:"not null;index"`
	ClassId     string    `gorm:"type:varchar(255);index"`
	StartTime   time.Time `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`

	Classroom Classroom `gorm:"foreignKey:ClassId;references:ClassId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Room      Room      `gorm:"foreignKey:RoomId;references:RoomId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	User      User      `gorm:"foreignKey:UserId;references:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
