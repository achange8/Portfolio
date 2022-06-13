package handler

import (
	"net/http"

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
	var boards []module.BOARD

	switch s_type {
	//title + content
	case "title_content":
		db := db.Connect()
		db.Where("TITLE LIKE ?", "%"+s_keyword+"%").Or("CONTENT LIKE ?", "%"+s_keyword+"%").Find(&boards)
		//should paging 1,2,...
		return c.JSON(http.StatusOK, boards)

	case "title":
		db := db.Connect()
		db.Where("TITLE LIKE ?", "%"+s_keyword+"%").Find(&boards)
		return c.JSON(http.StatusOK, boards)
	case "content":
		db := db.Connect()
		db.Where("CONTENT LIKE ?", "%"+s_keyword+"%").Find(&boards)
		return c.JSON(http.StatusOK, boards)
	case "witer":
		db := db.Connect()
		db.Where("WITER LIKE ?", "%"+s_keyword+"%").Find(&boards)
		return c.JSON(http.StatusOK, boards)
	}
	return c.JSON(http.StatusBadRequest, "bad request")
}
