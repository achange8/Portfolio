package handler

import (
	"net/http"

	db "github.com/achange8/Portfolio/DB"
	"github.com/achange8/Portfolio/module"
	"github.com/labstack/echo"
)

//test to get users
func GetAllUsers(c echo.Context) error {
	var list []module.User
	db := db.Connect()
	db.Table("users").Find(&list)
	return c.JSON(http.StatusOK, list)
}
