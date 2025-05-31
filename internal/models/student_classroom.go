package model

import (
	"time"
)

type StudentClass struct {
	ClassId   string    `gorm:"type:varchar(255);primaryKey;not null;index" json:"classID"`
	UserId    string    `gorm:"type:varchar(255);primaryKey;not null;index" json:"userID"`
	State     string    `gorm:"type:enum('joined','waiting')" json:"state"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	JoinedAt  time.Time `gorm:"autoUpdateTime" json:"joinedAt"`

	User      User      `gorm:"foreignKey:UserId;references:UserId;" json:"User"`
	Classroom Classroom `gorm:"foreignKey:ClassId;references:ClassId;" json:"Classroom"`
}
