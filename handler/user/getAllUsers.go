package user

import (
	"fmt"
	"net/http"

	database "github.com/achange8/Portfolio/DB"
	"github.com/achange8/Portfolio/module"
	"github.com/labstack/echo"
)

//test to get users
func GetAllUsers(c echo.Context) error {
	var list []module.User
	db := database.DB
	result := db.Table("users").Find(&list)
	fmt.Println(result.RowsAffected)
	return c.JSON(http.StatusOK, list)
}
