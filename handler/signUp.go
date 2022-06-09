package handler

import (
	"net/http"
	"regexp"

	db "github.com/achange8/Portfolio/DB"
	"github.com/achange8/Portfolio/module"
	"github.com/labstack/echo"
)

func SignUp(c echo.Context) error {
	db := db.Connect()
	user := new(module.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "bad request",
		})
	}

	savepassword := user.Password
	saveEmail := user.Email
	validID := regexp.MustCompile(`[ !@#$%^&*(),.?\":{}|<>]`)
	if validID.MatchString(user.Id) {
		return c.JSON(http.StatusBadRequest, "plz write A-Za-z0-9")
	}
	id := db.Raw("SELECT * FROM users WHERE id = ?", user.Id).Scan(&user)
	if id.RowsAffected != 0 {
		return c.JSON(http.StatusConflict, "ID already exists!")
	}
	validEmail, _ := regexp.Compile("^[_A-Za-z0-9+-.]+@[a-z0-9-]+(\\.[a-z0-9-]+)*(\\.[a-z]{2,4})$")
	if !validEmail.MatchString(saveEmail) {
		return c.JSON(http.StatusBadRequest, "plz write right email!")
	}
	email := db.Raw("SELECT * FROM users WHERE email = ?", saveEmail).Scan(&user)
	if email.RowsAffected != 0 {
		return c.JSON(http.StatusConflict, "Email already exists!")
	}
	if savepassword == "" {
		return c.JSON(http.StatusBadRequest, "bad password")
	}
	hashPW, err := module.HashPassword(savepassword)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "fail hashing PW")
	}
	user.Password = hashPW

	if err := db.Create(&user); err.Error != nil {
		return c.JSON(http.StatusBadRequest, "failed Sign Up")
	}
	return c.JSON(http.StatusOK, user)
}
