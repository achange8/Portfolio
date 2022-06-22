package module

import (
	database "github.com/achange8/Portfolio/DB"
)

//post
func Duplicate(id string) bool {
	db := database.DB
	user := new(User)
	a := db.Find(&user, "id = ?", id)

	return a.RowsAffected == 0
}
