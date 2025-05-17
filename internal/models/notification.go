package model

import (
	"time"
)

type Notification struct {
	NotificationId string `gorm:"type:varchar(255);primaryKey;not null;index"`
	UserId         string `gorm:"type:varchar(255);not null;index"`
	ClassId        string `gorm:"type:varchar(255);index"`
	Description    string
	Type           string    `gorm:"type:enum('success','warning','info')"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`

	Classroom Classroom `gorm:"foreignKey:ClassId;references:ClassId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	User      User      `gorm:"foreignKey:UserId;references:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
