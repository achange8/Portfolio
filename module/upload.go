package module

import (
	"io"
	"os"

	"github.com/labstack/echo"
)

func upload(c echo.Context) error {

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	for _, file := range form.File {
		// Source
		src, err := file[0].Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// Destination
		dst, err := os.Create(file[0].Filename)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}
	}
	return nil
}
