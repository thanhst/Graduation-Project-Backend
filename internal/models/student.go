package model

type Student struct {
	ID     uint   `gorm:"primaryKey;autoIncrement"`
	UserId string `gorm:"type:varchar(255);unique;index"`

	User          User           `gorm:"foreignKey:UserId;references:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	StudentClasss []StudentClass `gorm:"foreignKey:UserId;references:UserId"`
}
