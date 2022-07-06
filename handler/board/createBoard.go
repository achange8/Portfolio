package board

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	database "github.com/achange8/Portfolio/DB"
	"github.com/achange8/Portfolio/module"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

//check token in middleware, make token
//method : POST ,form-data, "files","data".
//data in content, title.
func CreateBoard(c echo.Context) error {
	cookie, err := c.Cookie("accessCookie")
	if err != nil {
		cookie, err = c.Cookie("RefreCookie")
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "no cookies sign in again")
		}
	}
	rawtoken := cookie.Value
	token, err := jwt.Parse(rawtoken, nil)
	if err == nil {
		return err
	}
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["files"]
	jsondata := form.Value["data"]
	var board module.BOARD
	err = json.Unmarshal([]byte(jsondata[0]), &board)
	if err != nil {
		fmt.Println(err)
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	board.WRITER = claims["jti"].(string)

	database.DB.Create(&board)

	for _, file := range files {
		// Source
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()
		contentType := file.Header.Get("Content-Type")
		if err != nil {
			return err
		}
		if contentType != "image/png" && contentType != "image/jpg" && contentType != "image/gif" {
			fmt.Printf("Unacceptable file type = %s\n", contentType)
			continue
		}

		// Destination
		dirpath := fmt.Sprintf("%d", board.NUM)
		dirname := "./uploads/" + dirpath

		os.MkdirAll(dirname, 0777)
		filepath := fmt.Sprintf("%s/%s", dirname, file.Filename)
		dst, err := os.Create(filepath)
		if err != nil {
			return err
		}
		//
		//todos : save filepath in db table
		//
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}
	}

	return c.JSON(http.StatusOK, board)
}
