package user

import (
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

//method GET
//return logged in user id if not logged in, return err
//JWT check process
func Usercheck(c echo.Context) error {
	envERR := godotenv.Load("db.env")
	if envERR != nil {
		log.Println("Could not load .env file")
		os.Exit(1)
	}
	cookie, err := c.Cookie("accessCookie")
	if err != nil {
		cookie, err = c.Cookie("RefreCookie")
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "no cookies sign in again")
		}
	}
	//auth token
	rawtoken := cookie.Value
	claims := &jwt.StandardClaims{}
	_, err = jwt.ParseWithClaims(rawtoken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("key")), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return c.JSON(http.StatusUnauthorized, "signature is invalid")
		}
		return c.JSON(http.StatusUnauthorized, "Authorization failed")
	}
	username := claims.Id
	if username == "" {
		return c.JSON(http.StatusUnauthorized, "no data")
	}
	return c.JSON(http.StatusOK, username)
}
