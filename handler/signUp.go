package handler

import (
	"net/http"

	db "github.com/achange8/Portfolio/DB"
	"github.com/achange8/Portfolio/module"
	"github.com/labstack/echo"
)

func SignUp(c echo.Context) error {
	db := db.Connect()
	user := new(module.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "bad request",
		})
	}

	savepassword := user.Password

	id := db.Find(user, "id=?", user.Id)
	if id.RowsAffected != 0 {
		return c.JSON(http.StatusForbidden, "ID or Email already exists!")
	}
	email := db.Find(user, "email=?", user.Email)

	if email.RowsAffected != 0 {
		return c.JSON(http.StatusForbidden, "ID or Email already exists!")
	}

	return c.JSON(http.StatusAccepted, map[string]string{"test done!": savepassword})
}
