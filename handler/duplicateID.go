package handler

import (
	"net/http"

	db "github.com/achange8/Portfolio/DB"
	"github.com/achange8/Portfolio/module"
	"github.com/labstack/echo"
)

//check id
func DuplCheckID(c echo.Context) error {
	db := db.Connect()
	user := new(module.User)
	err := c.Bind(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "failed bind request",
		})
	}
	id := db.Raw("SELECT * FROM users WHERE id = ?", user.Id).Scan(&user)
	if id.RowsAffected != 0 {
		return c.JSON(http.StatusConflict, "ID already exists!")
	}
	return c.JSON(http.StatusOK, "duplicate check done!")
}
