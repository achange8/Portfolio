package handler

import (
	"net/http"
	"strings"

	db "github.com/achange8/Portfolio/DB"
	"github.com/achange8/Portfolio/module"
	"github.com/labstack/echo"
)

//method : POST
//api:  /read/?id=num
func ReadBoard(c echo.Context) error {
	id := c.QueryParam("id")
	board := new(module.BOARD)
	viewcookie, err := c.Cookie("viewCookie")
	if err != nil {
		viewcookie = module.CreateViewCookie(id)
		c.SetCookie(viewcookie)
		db := db.Connect()
		db.Raw("UPDATE boards SET hi_tcount = ? WHERE NUM = ?", +1, id).Scan(board)
		return c.JSON(http.StatusOK, "make new cookie, view +1")
	}
	viewdata := strings.Split(viewcookie.Value, ",")
	result := check(viewdata, id)
	if result {
		newview := viewcookie.Value + "," + id
		viewcookie = module.CreateViewCookie(newview)
		c.SetCookie(viewcookie)
		db := db.Connect()
		db.Raw("UPDATE boards SET hi_tcount = ? WHERE NUM = ?", +1, id).Scan(board)
		return c.JSON(http.StatusOK, "reset cookie, view +1")
	}
	return c.JSON(http.StatusOK, "view +0")
}

func check(view []string, id string) bool {
	for i := range view {
		if view[i] == id {
			return false
		}
	}
	return true
}
