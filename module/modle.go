package module

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	P_id     int
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Refresh struct {
	Id       string
	Reftoken string
}

type ChangePWform struct {
	Oldpw string
	Newpw string
}

type ImgSrc struct {
	Imgname string
	Url     string
}
type Images struct {
	ImgSrc []ImgSrc
}

type BOARD struct {
	NUM       int    `gorm:"primaryKey"`
	TITLE     string `json:"TITLE"`
	WRITER    string `json:"WRITER"`
	CONTENT   string `json:"CONTENT"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"->:false;<-:create"`
	HiTCOUNT  int
}

type ListBoard struct {
	Board    []BOARD
	Lastpage int
}
