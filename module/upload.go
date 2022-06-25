package module

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func Upload(c echo.Context) error {

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["files"]
	jsondata := form.Value["data"]
	var board User
	err = json.Unmarshal([]byte(jsondata[0]), &board)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		// Source
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		contentType, err := GetFileContentType(src)
		if err != nil {
			return err
		}
		if contentType != "image/png" && contentType != "image/jpg" && contentType != "image/gif" {
			fmt.Printf("Unacceptable file type = %s\n", contentType)
			continue
		}
		// Destination
		dirname := "./uploads"
		os.MkdirAll(dirname, 0777)
		filepath := fmt.Sprintf("%s/%s", dirname, file.Filename)
		dst, err := os.Create(filepath)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}
	}

	return c.JSON(http.StatusOK, board)
}
