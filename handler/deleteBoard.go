package handler

import (
	"fmt"
	"net/http"
	"os"

	db "github.com/achange8/Portfolio/DB"
	"github.com/achange8/Portfolio/module"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

//if match jwt.id,db writer, delet board in db where board.num , delete uploads folder name %d,"num"
//method : board/delete/?num= , query parse "num"
func DeleteBoard(c echo.Context) error {
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
	db := db.Connect()
	db.Find(&board, "NUM = ?", num).Scan(board)
	if board.WRITER != writer {
		return c.JSON(http.StatusUnauthorized, "only writer can delete")
	}
	db.Where("NUM = ?", num).Delete(&board)
	path := fmt.Sprintf("./uploads/%d", board.NUM)
	err = os.RemoveAll(path)
	if err != nil {
		fmt.Println(err)
	}

	return c.JSON(http.StatusOK, "delete board done")
}
