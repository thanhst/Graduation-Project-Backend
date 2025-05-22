package model

import (
	"time"
)

type Scheduler struct {
	SchedulerId string    `gorm:"type:varchar(255);primaryKey;not null;index" json:"schedulerID"`
	RoomId      string    `gorm:"type:varchar(255);not null;index" json:"roomID"`
	UserId      string    `gorm:"not null;index" json:"userID"`
	ClassId     *string   `gorm:"type:varchar(255);index" json:"classID,omitempty"`
	StartTime   time.Time `gorm:"not null" json:"startTime"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updatedAt"`

	Classroom Classroom `gorm:"foreignKey:ClassId;references:ClassId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Room      Room      `gorm:"foreignKey:RoomId;references:RoomId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	User      User      `gorm:"foreignKey:UserId;references:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
