package model

import (
	"time"
)

type Account struct {
	AccountId   string     `json:"accountID" gorm:"type:varchar(255);primaryKey;unique;not null;index"`
	UserId      string     `json:"userID" gorm:"not null;type:varchar(255);index"`
	Email       string     `json:"email" gorm:"not null;"`
	Password    string     `json:"password" gorm:"not null"`
	Role        string     `json:"role" gorm:"type:enum('admin','teacher','student'); default:'student'"`
	Status      string     `json:"status" gorm:"type:enum('online','offline');default:'offline'"`
	LastLogin   *time.Time `json:"lastLogin,omitempty"`
	LoginMethod string     `json:"loginMethod" gorm:"type:enum('google','github','email');default:'email'"`
	CreatedAt   time.Time  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `json:"updatedAt" gorm:"autoUpdateTime"`

	User User `gorm:"foreignKey:UserId;references:UserId;"`
}
