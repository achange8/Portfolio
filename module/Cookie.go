package module

import (
	"net/http"
	"time"
)

func CreateAccCookie(userId string, ACtoken string) *http.Cookie {
	accessCookie := new(http.Cookie)
	accessCookie.Name = "accessCookie"
	accessCookie.Value = ACtoken
	accessCookie.Expires = time.Now().Add(30 * time.Minute)
	accessCookie.Path = "/"
	accessCookie.HttpOnly = true
	return accessCookie
}

func CreateRefreCookie(userId string, RefreToken string) *http.Cookie {
	refreshCookie := new(http.Cookie)
	refreshCookie.Name = "RefreCookie"
	refreshCookie.Value = RefreToken
	refreshCookie.Expires = time.Now().Add(24 * 30 * time.Hour)
	refreshCookie.Path = "/"
	refreshCookie.HttpOnly = true
	return refreshCookie
}

func LogOutAccCookie() *http.Cookie {
	accessCookie := new(http.Cookie)
	accessCookie.Name = "accessCookie"
	accessCookie.Value = ""
	accessCookie.Expires = time.Now().Add(-time.Hour)
	accessCookie.Path = "/"
	accessCookie.HttpOnly = true
	return accessCookie
}

func LogOutRefreCookie() *http.Cookie {
	refreshCookie := new(http.Cookie)
	refreshCookie.Name = "RefreCookie"
	refreshCookie.Value = ""
	refreshCookie.Expires = time.Now().Add(-time.Hour)
	refreshCookie.Path = "/"
	refreshCookie.HttpOnly = true
	return refreshCookie
}
