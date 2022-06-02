package handler

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig = oauth2.Config{
	RedirectURL:  "http://localhost:8082/auth/goolge/login",
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_SECRET_KEY"),
	Scopes:       []string{"https://www.googleapis.com/auth/contacts"},
	Endpoint:     google.Endpoint,
}

func GoogleLogin(c echo.Context) error {
	url := googleOauthConfig.AuthCodeURL()
	http.Redirect(w, r, http.StatusTemporaryRedirect)
	return nil
}
