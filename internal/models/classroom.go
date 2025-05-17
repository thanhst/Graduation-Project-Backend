package model

import (
	"time"
)

type Classroom struct {
	ClassId     string    `gorm:"type:varchar(255);primaryKey;unique;not null;index"`
	ClassName   string    `gorm:"not null;"`
	UserCreated string    `gorm:"not null;type:varchar(255);index"`
	Description string    `gorm:"not null"`
	Url         string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`

	Teacher Teacher `gorm:"foreignKey:UserCreated;references:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
