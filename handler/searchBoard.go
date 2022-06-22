package handler

import (
	"net/http"
	"strconv"

	database "github.com/achange8/Portfolio/DB"
	"github.com/achange8/Portfolio/module"
	"github.com/labstack/echo"
)

//search type , keywords in db and return contents [] /done
//s_type case : title+content, title, content, writer /done
//paging  	/done
// db connecting 2times
func SearchBoard(c echo.Context) error {
	s_type := c.QueryParam("type")
	s_keyword := c.QueryParam("s_keyword")
	page := c.QueryParam("page")
	if page == "" {
		page = "1"
	}

	n, err := strconv.Atoi(page)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}
	data := make([]module.ListBoard, 1)
	switch s_type {
	//title + content
	case "title_content":
		var count int64
		db := database.DB
		db.Model(&data[0].Board).Where("TITLE LIKE ?", "%"+s_keyword+"%").Or("CONTENT LIKE ?", "%"+s_keyword+"%").Count(&count)
		pages := int(count) / 10
		offset := (n - 1) * 10
		db.Where("TITLE LIKE ?", "%"+s_keyword+"%").Or("CONTENT LIKE ?", "%"+s_keyword+"%").Offset(offset).Limit(offset + 10).Order("NUM desc").Find(&data[0].Board)
		lastpage := pages + 1
		data[0].Lastpage = lastpage

		return c.JSON(http.StatusOK, data)

	case "title":
		db := database.DB
		db.Where("TITLE LIKE ?", "%"+s_keyword+"%").Find(&data[0].Board)
		return c.JSON(http.StatusOK, data[0].Board)
	case "content":
		db := database.DB
		db.Where("CONTENT LIKE ?", "%"+s_keyword+"%").Find(&data[0].Board)
		return c.JSON(http.StatusOK, data[0].Board)
	case "witer":
		db := database.DB
		db.Where("WITER LIKE ?", "%"+s_keyword+"%").Find(&data[0].Board)
		return c.JSON(http.StatusOK, data[0].Board)
	}
	return c.JSON(http.StatusBadRequest, "bad request")
}
