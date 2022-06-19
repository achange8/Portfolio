package handler

import (
	"fmt"
	"log"

	"github.com/labstack/echo"
)

func DownLoadFile(c echo.Context) error {
	num := c.QueryParam("num")
	filename := c.QueryParam("name")
	fmt.Printf("num = %s, fname = %s", num, filename)
	peth := fmt.Sprintf("./uploads/%s/%s", num, filename)
	return c.File(peth)
}

func LoadFile(c echo.Context) error {
	num := c.QueryParam("num")
	filename := c.QueryParam("name")
	log.Printf("num = %s, fname = %s", num, filename)
	peth := fmt.Sprintf("./uploads/%s/%s", num, filename)
	return c.Attachment(peth, filename)
}
