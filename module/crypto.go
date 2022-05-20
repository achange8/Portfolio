package module

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPW_Hash(hashval, userPW string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashval), []byte(userPW))
	if err != nil {
		return false
	}
	return true
}
