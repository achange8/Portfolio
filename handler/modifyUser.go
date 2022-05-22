package handler

import (
	"fmt"
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
	saveEmail := user.Email
	saveId := user.Id
	savePW := user.Password

	id := db.Raw("SELECT * FROM users WHERE id = ?", user.Id).Scan(&user)
	if id.RowsAffected != 0 {
		fmt.Println(user)
		return c.JSON(http.StatusConflict, "ID already exists!")
	}

	db.Model(&user).Where("email = ?", saveEmail).
		Updates(map[string]interface{}{"id": saveId, "password": savePW})

	return c.JSON(http.StatusOK, user)
}
