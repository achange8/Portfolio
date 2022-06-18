package module

import (
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

	for _, file := range files {
		// Source
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// Destination
		dirname := "./uploads"
		os.MkdirAll(dirname, 0777)
		filepath := fmt.Sprintf("%s/upload-%s", dirname, file.Filename)
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
	return c.JSON(http.StatusOK, "done")
}
