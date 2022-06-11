package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

//search type , keywords in db and return contents []
//s_type case : title+content, title, content, writer
func SearchBoard(c echo.Context) error {
	s_type := c.QueryParam("type")
	s_keyword := c.QueryParam("s_keyword")

	switch s_type {
	case "title_content":
		return c.JSON(http.StatusOK, map[string]string{
			"s_type":    "title_content",
			"s_keyword": s_keyword,
		})

	case "title":
		return c.JSON(http.StatusOK, map[string]string{
			"s_type":    "title",
			"s_keyword": s_keyword,
		})
	case "content":
		return c.JSON(http.StatusOK, map[string]string{
			"s_type":    "content",
			"s_keyword": s_keyword,
		})
	case "witer":
		return c.JSON(http.StatusOK, map[string]string{
			"s_type":    "witer",
			"s_keyword": s_keyword,
		})
	}
	return c.JSON(http.StatusBadRequest, "bad request")
}
