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
	db := database.DB
	var count int64
	switch s_type {
	//title + content
	case "title_content":
		db.Model(&data[0].Board).Where("TITLE LIKE ?", "%"+s_keyword+"%").Or("CONTENT LIKE ?", "%"+s_keyword+"%").Count(&count)
		data[0].Lastpage = int(count)/10 + 1
		offset := (n - 1) * 10
		db.Where("TITLE LIKE ?", "%"+s_keyword+"%").Or("CONTENT LIKE ?", "%"+s_keyword+"%").Offset(offset).Limit(offset + 10).Order("NUM desc").Find(&data[0].Board)

		return c.JSON(http.StatusOK, data)

	case "title":
		db.Model(&data[0].Board).Where("TITLE LIKE ?", "%"+s_keyword+"%").Count(&count)
		data[0].Lastpage = int(count)/10 + 1
		offset := (n - 1) * 10
		db.Where("TITLE LIKE ?", "%"+s_keyword+"%").Offset(offset).Limit(offset + 10).Order("NUM desc").Find(&data[0].Board)

		return c.JSON(http.StatusOK, data)

	case "content":
		db.Model(&data[0].Board).Where("CONTENT LIKE ?", "%"+s_keyword+"%").Count(&count)
		data[0].Lastpage = int(count)/10 + 1
		offset := (n - 1) * 10
		db.Where("CONTENT LIKE ?", "%"+s_keyword+"%").Offset(offset).Limit(offset + 10).Order("NUM desc").Find(&data[0].Board)

		return c.JSON(http.StatusOK, data)

	case "witer":
		db.Model(&data[0].Board).Where("WITER LIKE ?", "%"+s_keyword+"%").Count(&count)
		data[0].Lastpage = int(count)/10 + 1
		offset := (n - 1) * 10
		db.Where("WITER LIKE ?", "%"+s_keyword+"%").Offset(offset).Limit(offset + 10).Order("NUM desc").Find(&data[0].Board)

		return c.JSON(http.StatusOK, data)
	}

	return c.JSON(http.StatusBadRequest, "bad request")
}
