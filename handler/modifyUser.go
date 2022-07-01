package handler

import (
	"log"
	"net/http"
	"os"

	database "github.com/achange8/Portfolio/DB"
	"github.com/achange8/Portfolio/module"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

//change user ID
func ModifyID(c echo.Context) error {
	db := database.DB
	user := new(module.User)
	err := c.Bind(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "failed bind request",
		})
	}
	changeId := user.Id
	envERR := godotenv.Load("db.env")
	if envERR != nil {
		log.Println("Could not load .env file")
		os.Exit(1)
	}

	cookie, err := c.Cookie("RefreCookie")
	if err != nil || cookie.Value == "" {
		log.Printf("not logged in user")
		return c.JSON(http.StatusUnauthorized, "login plz")
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

	db.Raw("UPDATE db.users SET id = ? WHERE id = ?", changeId, refreshClaims.Id).Scan(&user)

	ACCookie := module.LogOutAccCookie()
	RFCookie := module.LogOutRefreCookie()
	c.SetCookie(ACCookie)
	c.SetCookie(RFCookie)

	return c.JSON(http.StatusOK, "ID change done, plz login again")
}

//change PW
func ModifyPW(c echo.Context) error {
	db := database.DB
	PWform := new(module.ChangePWform) //include change PW
	err := c.Bind(PWform)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "failed bind request",
		})
	}
	envERR := godotenv.Load("db.env")
	if envERR != nil {
		log.Println("Could not load .env file")
		os.Exit(1)
	}

	cookie, err := c.Cookie("RefreCookie")
	if err != nil || cookie.Value == "" {
		log.Printf("not logged in user")
		return c.JSON(http.StatusUnauthorized, "login plz")
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
	user := new(module.User)
	db.Find(&user, "id = ?", refreshClaims.Id).Scan(user)
	checkpw := module.CheckPW_Hash(user.Password, PWform.Oldpw)
	if !checkpw {
		return c.JSON(http.StatusUnauthorized, "wrong password")
	}
	hashNewPW, _ := module.HashPassword(PWform.Newpw)
	db.Raw("UPDATE users SET password = ? WHERE password = ?", hashNewPW, user.Password)

	ACCookie := module.LogOutAccCookie()
	RFCookie := module.LogOutRefreCookie()
	c.SetCookie(ACCookie)
	c.SetCookie(RFCookie)
	return c.JSON(http.StatusOK, "PW change Done,plz login again")
}
