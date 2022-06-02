package handler

import (
	"net/http"

	db "github.com/achange8/Portfolio/DB"
	"github.com/achange8/Portfolio/module"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

func UserDelete(c echo.Context) error {
	user := new(module.User)
	err := c.Bind(user)
	saveEmail := user.Email
	if err != nil {
		return c.JSON(http.StatusBadRequest, "failed bind context")
	}
	cookie, err := c.Cookie("accessCookie")
	if err != nil {
		cookie, err = c.Cookie("RefreCookie")
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "no cookies sign in again")
		}
	}
	rawtoken := cookie.Value
	token, _ := jwt.Parse(rawtoken, nil)
	claims, _ := token.Claims.(jwt.MapClaims)
	userid := claims["jti"].(string)
	db := db.Connect()
	db.Find(&user, "id=?", userid).Scan(user)
	if user.Email != saveEmail {
		return c.JSON(http.StatusUnauthorized, "not same email")
	}
	//DELETE from users where id = userid;
	db.Where("id=?", userid).Delete(&user)
	logoutACcookie := module.LogOutAccCookie()
	logoutRFcookie := module.LogOutRefreCookie()
	c.SetCookie(logoutACcookie)
	c.SetCookie(logoutRFcookie)

	return c.JSON(http.StatusOK, "delete done")
}
