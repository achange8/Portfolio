package module

import db "github.com/achange8/Portfolio/DB"

//post
func Duplicate(id string) bool {
	db := db.Connect()
	user := new(User)
	a := db.Find(&user, "id = ?", id)

	return a.RowsAffected == 0
}
