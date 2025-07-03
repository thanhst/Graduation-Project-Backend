package model

import (
	"time"
)

type Notification struct {
	NotificationId string    `gorm:"type:varchar(255);primaryKey;not null;index" json:"id"`
	UserId         string    `gorm:"type:varchar(255);not null;index" json:"userID"`
	ClassId        *string   `gorm:"type:varchar(255);index" json:"classID"`
	SchedulerId    *string   `gorm:"type:varchar(255);index" json:"schedulerID"`
	Description    string    `json:"descriptions"`
	Type           string    `gorm:"type:enum('success','warning','info')" json:"type"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"createdAt"`

	Classroom Classroom `gorm:"foreignKey:ClassId;references:ClassId;" json:"Classroom"`
	User      User      `gorm:"foreignKey:UserId;references:UserId;" json:"User"`
	Scheduler Scheduler `gorm:"foreignKey:SchedulerId;references:SchedulerId;" json:"Scheduler"`
}
