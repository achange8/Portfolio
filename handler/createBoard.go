package handler

import (
	"net/http"

	db "github.com/achange8/Portfolio/DB"
	"github.com/achange8/Portfolio/module"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

//check token in middleware, make token
//click write button, work this api
//method : POST
func CreateBoard(c echo.Context) error {
	cookie, err := c.Cookie("accessCookie")
	if err != nil {
		cookie, err = c.Cookie("RefreCookie")
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "no cookies sign in again")
		}
	}
	rawtoken := cookie.Value
	token, err := jwt.Parse(rawtoken, nil)
	if err == nil {
		return err
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	board := new(module.BOARD)
	board.WRITER = claims["jti"].(string)
	err = c.Bind(board)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "failed bind board")
	}
	db := db.Connect()
	db.Create(&board)
	return c.JSON(http.StatusOK, "write board done")
}
