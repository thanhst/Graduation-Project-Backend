package model

import (
	"time"
)

type Emotion struct {
	Id        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	RoomId    string    `gorm:"type:varchar(255);index;not null" json:"roomID"`
	UserId    string    `gorm:"type:varchar(255);index;not null" json:"userID"`
	Emotion   string    `gorm:"type:enum('Happy','Sad','Neutral','Fear','Surprise')" json:"emotion"`
	CreatedAt time.Time `gorm:"autoCreateTime;index" json:"createdAt"`

	Room Room `gorm:"foreignKey:RoomId;references:RoomId;"`
	User User `gorm:"foreignKey:UserId;references:UserId;"`
}
