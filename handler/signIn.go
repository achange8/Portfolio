package handler

import (
	"log"
	"net/http"
	"regexp"

	db "github.com/achange8/Portfolio/DB"
	"github.com/achange8/Portfolio/module"
	"github.com/labstack/echo"
)

//check id,pw and give ac,rf token
//POST localhost/signin
func SignIn(c echo.Context) error {
	db := db.Connect()
	user := new(module.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "bad request",
		})
	}
	savepassword := user.Password
	validID := regexp.MustCompile(`[ !@#$%^&*(),.?\":{}|<>]`)
	if validID.MatchString(user.Id) {
		return c.JSON(http.StatusBadRequest, "plz write A-Za-z0-9")
	}
	id := db.Raw("SELECT * FROM users WHERE id = ?", user.Id).Scan(&user)
	if id.RowsAffected == 0 {
		return c.JSON(http.StatusUnauthorized, "ID is worng")
	}
	if savepassword == " " {
		return c.JSON(http.StatusUnauthorized, "bad password")
	}
	checkpw := module.CheckPW_Hash(user.Password, savepassword)
	if !checkpw {
		c.JSON(http.StatusUnauthorized, "wrong password")
	}
	//give JWT cookie code
	accesstoken, err := module.CreateAccToken(user.Id)
	if err != nil {
		log.Println("err creating access token!")
	}
	refreshtoken, err := module.CreateRefreshToken(user.Id)
	if err != nil {
		log.Println("err creating refresh token!")
	}
	accessCookie := module.CreateAccCookie(user.Id, accesstoken)
	c.SetCookie(accessCookie)
	refreshCookie := module.CreateRefreCookie(user.Id, refreshtoken)
	c.SetCookie(refreshCookie)
	refresh := new(module.Refresh)
	id = db.Find(&refresh, "id=?", user.Id)
	if id.RowsAffected != 0 {
		//db.Model(&refresh).Where("id =?", user.Id).Update("reftoken", refreshtoken)
		db.Raw("UPDATE refreshes SET reftoken = ? WHERE id = ?", refreshtoken, user.Id).Scan(&user)

		// UPDATE refreshes SET `reftoken` = RefreshToken WHERE id = user.Id
	} else {
		refresh.Id = user.Id
		refresh.Reftoken = refreshtoken
		db.Create(&refresh)
	}

	return c.JSON(http.StatusOK,
		map[string]string{
			"message":      "ok",
			"accesstoken":  accesstoken,
			"refreshtoken": refreshtoken})
}
