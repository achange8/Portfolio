package handler

import (
	"net/http"
	"strings"

	db "github.com/achange8/Portfolio/DB"
	"github.com/achange8/Portfolio/module"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
		db.Model(&board).Clauses(clause.Returning{}).Where("NUM = ?", id).Update("hi_tcount", gorm.Expr("hi_tcount + ?", 1)).Scan(board)

		return c.JSON(http.StatusOK, board)
	}
	viewdata := strings.Split(viewcookie.Value, ",")
	result := check(viewdata, id)
	if result {
		newview := viewcookie.Value + "," + id
		viewcookie = module.CreateViewCookie(newview)
		c.SetCookie(viewcookie)
		db := db.Connect()
		db.Model(&board).Clauses(clause.Returning{}).Where("NUM = ?", id).Update("hi_tcount", gorm.Expr("hi_tcount + ?", 1)).Scan(board)

		return c.JSON(http.StatusOK, board)
	}
	db := db.Connect()
	db.Find(&board, "NUM =?", id).Scan(board)
	return c.JSON(http.StatusOK, board)
}

func check(view []string, id string) bool {
	for i := range view {
		if view[i] == id {
			return false
		}
	}
	return true
}
