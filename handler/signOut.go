package handler

import (
	"log"
	"net/http"
	"os"

	database "github.com/achange8/Portfolio/DB"
	"github.com/achange8/Portfolio/module"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

//method get localhost/signout
func SignOut(c echo.Context) error {
	cookie, err := c.Cookie("RefreCookie")
	if err != nil {
		return nil

	} else {
		refresh := new(module.Refresh)
		envERR := godotenv.Load("db.env")
		if envERR != nil {
			log.Println("Could not load db.env file")
			os.Exit(1)
		}
		token := cookie.Value
		refreshClaims := &jwt.StandardClaims{}
		_, err = jwt.ParseWithClaims(token, refreshClaims,
			func(t *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("key2")), nil
			})
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "jwt not allowed")
		}
		//delete reftoken in db
		db := database.DB
		result := db.Find(&refresh, "reftoken=?", token)
		if result.RowsAffected != 0 { //todo go signin point
			db.Where("reftoken = ?", token).Delete(&refresh)
		}
		//delete cookies (create - time cookie)
		accessCookie := module.LogOutAccCookie()
		c.SetCookie(accessCookie)
		RefreCookie := module.LogOutRefreCookie()
		c.SetCookie(RefreCookie)

		return c.JSON(http.StatusOK, "Logged out done")
	}
}

//del acckookies
