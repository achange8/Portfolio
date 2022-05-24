package module

import db "github.com/achange8/Portfolio/DB"

func Duplicate(id string) bool {
	db := db.Connect()
	user := new(User)
	a := db.Find(&user, "id = ?", id)
	if a.RowsAffected == 0 {
		return true
	}
	return false
}
