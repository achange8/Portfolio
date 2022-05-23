package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"

	db "github.com/achange8/Portfolio/DB"
	"github.com/achange8/Portfolio/module"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func Modifyuser(c echo.Context) error {
	db := db.Connect()
	user := new(module.User)
	err := c.Bind(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "failed bind request",
		})
	}
	saveEmail := user.Email
	saveId := user.Id
	savePW := user.Password
	envERR := godotenv.Load("db.env")
	if envERR != nil {
		log.Println("Could not load .env file")
		os.Exit(1)
	}

	cookie, err := c.Cookie("RefreCookie")
	if err != nil {
		return err
	}
	token := cookie.Value
	refreshClaims := &jwt.StandardClaims{}
	_, err = jwt.ParseWithClaims(token, refreshClaims,
		func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("key2")), nil
		})
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "jwt not allowed")
	}
	fmt.Println("parse done")
	id := db.Raw("SELECT * FROM users WHERE id = ?", user.Id).Scan(&user)
	if id.RowsAffected != 0 {
		return c.JSON(http.StatusConflict, "ID already exists!")
	} else {
		pw, _ := module.HashPassword(savePW)
		db.Model(&user).Where("email = ?", saveEmail).
			Updates(map[string]interface{}{"id": saveId, "password": pw})
	}
	fmt.Println(user)
	db.Raw("SELECT * FROM users WHERE id = ?", user.Id).Scan(&user)

	return c.JSON(http.StatusOK, user)
}
