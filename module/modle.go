package module

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	P_id     int `gorm:"autoIncrement`
	Id       string
	Email    string
	Password string
}

type Refresh struct {
	Id       string
	Reftoken string
}

type ChangePWform struct {
	Oldpw string
	Newpw string
}

type BOARD struct {
	NUM       int `gorm:"primaryKey"`
	TITLE     string
	WRITER    string
	CONTENT   string
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"->:false;<-:create"`
	HiTCOUNT  int
}
