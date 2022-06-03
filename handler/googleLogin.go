package handler

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig = oauth2.Config{
	RedirectURL:  "http://localhost:8082/auth/goolge/callback",
	ClientID:     "295221415874-mmh9djapdgnsl7i8neke5kec0984a9fm.apps.googleusercontent.com",
	ClientSecret: "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",
	Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile"},
	Endpoint: google.Endpoint,
}

func GoogleLogin(c echo.Context) error {
	state := oauthCookie(c)
	url := googleOauthConfig.AuthCodeURL(state)
	c.Redirect(http.StatusTemporaryRedirect, url)

	return nil
}

func oauthCookie(c echo.Context) string {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := new(http.Cookie)
	cookie.Name = "oauthcookie"
	cookie.Value = state
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
	return state
}

func GoogleCallBack(c echo.Context) error {
	cookie, err := c.Cookie("oauthcookie")
	if err != nil {
		return c.JSON(http.StatusBadRequest, "no cookie")
	}
	if c.FormValue("state") != cookie.Value {
		log.Println("invalid google oauth state cookie : ", cookie.Value)
		return c.JSON(http.StatusUnauthorized, "invailed state")
	}
	data, err := getGoogleUserInfo(c.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		c.Redirect(http.StatusUnauthorized, "/")
	}
	return c.JSON(http.StatusOK, string(data))

}

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func getGoogleUserInfo(code string) ([]byte, error) {
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange %s\n", err.Error())
	}
	res, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to get user info %s\n", err.Error())
	}
	return io.ReadAll(res.Body)
}
