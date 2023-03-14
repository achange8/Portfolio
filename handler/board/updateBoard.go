package board

import (
	"fmt"
	"net/http"

	database "github.com/achange8/Portfolio/DB"
	"github.com/achange8/Portfolio/module"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

// method : POST
// : /board/modify/?num=
// db connect time : 2 times
func UpdateBoard(c echo.Context) error {
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
		fmt.Println("err", err)
	}
	claims := token.Claims.(jwt.MapClaims)
	writer := claims["jti"].(string)
	num := c.QueryParam("num")
	board := new(module.BOARD)
	boarderr := c.Bind(board)
	if boarderr != nil {
		return c.JSON(http.StatusBadRequest, "failed bind")
	}
	savetitle := board.TITLE
	savecontent := board.CONTENT
	db := database.DB
	db.Find(&board, "NUM = ?", num).Scan(board)
	if board.WRITER != writer {
		return c.JSON(http.StatusUnauthorized, "only writer can")
	}
	db.Model(&board).Where("NUM = ?", num).Updates(module.BOARD{TITLE: savetitle, CONTENT: savecontent})
	return c.JSON(http.StatusOK, board)
}
