package handler

import (
	"net/http"
	"strings"

	"github.com/achange8/Portfolio/module"
	"github.com/labstack/echo"
)

//method : POST
//api:  /read/?id=num
func ReadBoard(c echo.Context) error {
	id := c.QueryParam("id")
	viewcookie, err := c.Cookie("viewCookie")
	if err != nil {
		viewcookie = module.CreateViewCookie(id)
		c.SetCookie(viewcookie)
		return c.JSON(http.StatusOK, "make new cookie")
	}
	viewdata := strings.Split(viewcookie.Value, ",")
	result := check(viewdata, id)
	if result {
		newview := viewcookie.Value + "," + id
		viewcookie = module.CreateViewCookie(newview)
		c.SetCookie(viewcookie)
		return c.JSON(http.StatusOK, map[string]string{
			"type":  "reset view cookie",
			"value": newview,
		})
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
