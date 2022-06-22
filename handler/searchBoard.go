package handler

import (
	"net/http"
	"strconv"

	db "github.com/achange8/Portfolio/DB"
	"github.com/achange8/Portfolio/module"
	"github.com/labstack/echo"
)

//search type , keywords in db and return contents [] /done
//s_type case : title+content, title, content, writer /done
//paging  	/not done
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
		db := db.Connect()
		var count int64
		db.Model(&data[0].Board).Where("TITLE LIKE ?", "%"+s_keyword+"%").Or("CONTENT LIKE ?", "%"+s_keyword+"%").Count(&count)
		total := count

		pages := int(total) / 10
		offset := (n - 1) * 10
		//limit := 10*(n-1) + end
		db.Where("TITLE LIKE ?", "%"+s_keyword+"%").Or("CONTENT LIKE ?", "%"+s_keyword+"%").Offset(offset).Limit(offset + 10).Order("NUM desc").Find(&data[0].Board)
		lastpage := pages + 1
		data[0].Lastpage = lastpage

		return c.JSON(http.StatusOK, data)

	case "title":
		db := db.Connect()
		db.Where("TITLE LIKE ?", "%"+s_keyword+"%").Find(&data[0].Board)
		return c.JSON(http.StatusOK, data[0].Board)
	case "content":
		db := db.Connect()
		db.Where("CONTENT LIKE ?", "%"+s_keyword+"%").Find(&data[0].Board)
		return c.JSON(http.StatusOK, data[0].Board)
	case "witer":
		db := db.Connect()
		db.Where("WITER LIKE ?", "%"+s_keyword+"%").Find(&data[0].Board)
		return c.JSON(http.StatusOK, data[0].Board)
	}
	return c.JSON(http.StatusBadRequest, "bad request")
}
