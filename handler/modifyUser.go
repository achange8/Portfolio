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

	return c.JSON(http.StatusOK, "ID change done")
}

//change PW
func ModifyPW(c echo.Context) error {
	db := db.Connect()
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

	//todos
	//새로운 구조체로 json 입력받아서 파싱처리 done
	// db에서 ref token 값, 입력받은 pw 해쉬값 비교
	//일치하면 users pw를 바꿀 pw로 변경
	user := new(module.User)
	db.Find(&user, "user = ?", refreshClaims.Id).Scan(user)
	checkpw := module.CheckPW_Hash(user.Password, PWform.OldPW)
	if !checkpw {
		return c.JSON(http.StatusUnauthorized, "wrong password")
	}
	hashNewPW, _ := module.HashPassword(PWform.NewPW)
	db.Raw("UPDATE users SET password = ? WHERE password = ?", hashNewPW, user.Password).Scan(&user)

	return c.JSON(http.StatusOK, "PW change Done")
}
