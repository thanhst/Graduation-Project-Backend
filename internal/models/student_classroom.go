package model

import (
	"time"
)

type StudentClass struct {
	ClassId   string    `gorm:"type:varchar(255);primaryKey;not null;index"`
	UserId    string    `gorm:"type:varchar(255);primaryKey;not null;index"`
	State     string    `gorm:"type:enum('joined','waiting')"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	JoinedAt  time.Time `gorm:"autoUpdateTime"`

	User      User      `gorm:"foreignKey:UserId;references:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Classroom Classroom `gorm:"foreignKey:ClassId;references:ClassId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
