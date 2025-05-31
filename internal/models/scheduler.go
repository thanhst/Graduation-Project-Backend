package model

import (
	"time"
)

type Scheduler struct {
	SchedulerId string    `gorm:"type:varchar(255);primaryKey;not null;index" json:"schedulerID"`
	RoomId      string    `gorm:"type:varchar(255);not null;index" json:"roomID"`
	UserId      string    `gorm:"not null;index" json:"userID"`
	ClassId     *string   `gorm:"type:varchar(255);" json:"classID,omitempty"`
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	StartTime   time.Time `gorm:"not null" json:"startTime"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updatedAt"`

	Classroom Classroom `gorm:"foreignKey:ClassId;references:ClassId;"`
	Room      Room      `gorm:"foreignKey:RoomId;references:RoomId;"`
	User      User      `gorm:"foreignKey:UserId;references:UserId;"`
}
