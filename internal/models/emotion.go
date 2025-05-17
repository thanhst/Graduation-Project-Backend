package model

import (
	"time"
)

type Emotion struct {
	Id        int       `gorm:"primaryKey;autoIncrement"`
	RoomId    string    `gorm:"type:varchar(255);index;not null"`
	UserId    string    `gorm:"type:varchar(255);index;not null"`
	Emotion   string    `gorm:"type:enum('Happy','Sad','Neutral','Fear','Surprise')"`
	CreatedAt time.Time `gorm:"autoCreateTime;index"`

	Room Room `gorm:"foreignKey:RoomId;references:RoomId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	User User `gorm:"foreignKey:UserId;references:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
