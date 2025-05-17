package model

type Teacher struct {
	ID     uint   `gorm:"primaryKey;autoIncrement"`
	UserId string `gorm:"type:varchar(255);index"`

	User       User        `gorm:"foreignKey:UserId;references:UserId;"`
	Classrooms []Classroom `gorm:"foreignKey:UserCreated;references:UserId"`
}
