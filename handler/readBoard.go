package handler

import (
	"net/http"

	db "github.com/achange8/Portfolio/DB"
	"github.com/achange8/Portfolio/module"
	"github.com/labstack/echo"
)

//method get
//api /readboard
func ReadBoard(c echo.Context) error {
	var list []module.BOARD
	db := db.Connect()
	db.Table("boards").Find(&list)
	return c.JSON(http.StatusOK, list)
}
