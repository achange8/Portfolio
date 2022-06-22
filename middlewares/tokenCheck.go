package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"os"

	database "github.com/achange8/Portfolio/DB"
	"github.com/achange8/Portfolio/module"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

//log in check ACToken
//if dont have ACToken, look up RefToken
//if have ref token, recreate ac token else return status unhorized
func TokenCheckMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		envERR := godotenv.Load("db.env")
		if envERR != nil {
			log.Println("Could not load .env file")
			os.Exit(1)
		}
		cookie, err := c.Cookie("accessCookie")
		//if have access token
		if err == nil {
			rawtoken := cookie.Value
			clamis := &jwt.StandardClaims{}
			_, err := jwt.ParseWithClaims(rawtoken, clamis, func(t *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("key")), nil
			})
			if err != nil {
				return c.JSON(http.StatusUnauthorized, "jwt not allowed")
			}
			return next(c)
		} else {
			//if dont have access token, check ref token
			cookie, err := c.Cookie("RefreCookie")
			if err != nil {
				c.JSON(http.StatusUnauthorized, "you dont have reftoken")
			}
			rawtoken := cookie.Value
			claims := &jwt.StandardClaims{}
			_, err = jwt.ParseWithClaims(rawtoken, claims,
				func(t *jwt.Token) (interface{}, error) {
					return []byte(os.Getenv("key2")), nil
				})
			if err != nil {
				return c.JSON(http.StatusUnauthorized, "jwt not allowed , log in again")
			}
			db := database.DB
			refresh := new(module.Refresh)
			db.Find(&refresh, "reftoken=?", rawtoken).Scan(&refresh)
			fmt.Println(refresh)
			if refresh.Id != claims.Id {
				return c.JSON(http.StatusUnauthorized, "Do signin again")
			} else {
				newtoken, _ := module.CreateAccToken(claims.Id)
				cookie := module.CreateAccCookie(claims.Id, newtoken)
				c.SetCookie(cookie)

				return next(c)
			}
		}
	}
}
