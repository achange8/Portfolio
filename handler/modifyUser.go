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

//change user ID
func ModifyID(c echo.Context) error {
	db := db.Connect()
	user := new(module.User)
	err := c.Bind(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "failed bind request",
		})
	}
	fmt.Println(user)
	changeId := user.Id
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
	if !module.Duplicate(changeId) {
		return c.JSON(http.StatusBadRequest, "ID already exists!")
	}

	db.Raw("UPDATE users SET id = ? WHERE id = ?", changeId, refreshClaims.Id).Scan(&user)

	return c.JSON(http.StatusOK, user)
}

//change PW
func ModifyPW(c echo.Context) error {
	db := db.Connect()
	user := new(module.User) //include change PW
	err := c.Bind(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "failed bind request",
		})
	}
	changePW := user.Password
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
	RealUserID := refreshClaims.Id
	db.Find(&user, "user = ?", refreshClaims.Id).Scan(user)
	db.Raw("UPDATE users SET PW = ? WHERE PW = ?", changePW, refreshClaims.Id).Scan(&user)

	return c.JSON(http.StatusOK, user)
}
