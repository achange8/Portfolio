package board

import (
	"net/http"
	"strconv"

	database "github.com/achange8/Portfolio/DB"
	"github.com/achange8/Portfolio/module"
	"github.com/labstack/echo"
)

//method get
//api /list?page=
func ListBoard(c echo.Context) error {
	page := c.QueryParam("page")
	if page == "" {
		page = "1"
	}
	var list []module.BOARD
	n, err := strconv.Atoi(page)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}
	result := database.DB.Table("boards").Order("NUM desc").Find(&list)
	total := result.RowsAffected
	println(total)
	pages := int(total) / 10
	end := n * 10
	if n == pages+1 {
		end = int(total) % 10
	}

	return c.JSON(http.StatusOK, list[(n-1)*10:10*(n-1)+end])
}
