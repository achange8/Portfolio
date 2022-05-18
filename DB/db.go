package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	envERR := godotenv.Load("db.env")
	if envERR != nil {
		fmt.Println("could not load env file !")
		os.Exit(1)
	}
	USER := os.Getenv("UserName")
	PASS := os.Getenv("PassWord")
	Protocol := os.Getenv("Protocal")
	DB_Name := os.Getenv("DBname")

	CONNECT := USER + ":" + PASS + "@" + Protocol + "/" + DB_Name + "?charset=utf8mb4&parseTime=True&loc=Asia%2FSeoul"
	db, err := gorm.Open(mysql.Open(CONNECT), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db
}
