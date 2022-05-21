package handler

import (
	"net/http"

	db "github.com/achange8/Portfolio/DB"
	"github.com/achange8/Portfolio/module"
	"github.com/labstack/echo"
)

func Modifyuser(c echo.Context) error {
	db := db.Connect()
	user := new(module.User)
	err := c.Bind(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "failed bind request",
		})
	}
	id := db.Find(user, "id=?", user.Id)
	if id.RowsAffected != 0 {
		return c.JSON(http.StatusBadRequest, "existing id")
	}
	db.Model(&user).Where("email = ?", user.Email).
		Updates(map[string]interface{}{"id": user.Id, "password": user.Password})

	return c.JSON(http.StatusOK, "user modify done!")
}
