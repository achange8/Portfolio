package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	envERR := godotenv.Load(".env")
	if envERR != nil {
		fmt.Println("could not load env file !")
		os.Exit(1)
	}
	USER := os.Getenv("UserName")
	PASS := os.Getenv("PassWord")
	Protocol := os.Getenv("Protocol")
	DB_Name := os.Getenv("DBname")

	CONNECT := USER + ":" + PASS + "@" + Protocol + "/" + DB_Name + "?charset=utf8mb4&parseTime=True&loc=Asia%2FSeoul"
	db, err := gorm.Open(mysql.Open(CONNECT), &gorm.Config{})
	if err != nil {
		println("err 1")
		panic(err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		println("2nd err ")
		println(err.Error())
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	DB = db
}
