package model

import (
	"time"
)

type Classroom struct {
	ClassId     string    `gorm:"type:varchar(255);primaryKey;unique;not null;index" json:"classID"`
	ClassName   string    `gorm:"not null;" json:"className"`
	UserCreated string    `gorm:"not null;type:varchar(255);index" json:"userCreated"`
	Description string    `gorm:"not null" json:"descriptions"`
	Url         string    `gorm:"not null" json:"url"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updatedAt"`

	Teacher        Teacher        `gorm:"foreignKey:UserCreated;references:UserId;"`
	User           User           `gorm:"foreignKey:UserCreated;references:UserId;"`
	StudentClasses []StudentClass `gorm:"foreignKey:ClassId;references:ClassId;"`
	Schedulers     []Scheduler    `gorm:"foreignKey:ClassId;references:ClassId"`
}
