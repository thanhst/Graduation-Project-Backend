package model

import (
	"time"
)

type Notification struct {
	NotificationId string    `gorm:"type:varchar(255);primaryKey;not null;index" json:"id"`
	UserId         string    `gorm:"type:varchar(255);not null;index" json:"userID"`
	ClassId        string    `gorm:"type:varchar(255);index" json:"classID"`
	Description    string    `json:"descriptions"`
	Type           string    `gorm:"type:enum('success','warning','info')" json:"type"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"createdAt"`

	Classroom Classroom `gorm:"foreignKey:ClassId;references:ClassId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	User      User      `gorm:"foreignKey:UserId;references:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
