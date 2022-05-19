package router

import (
	"github.com/achange8/Portfolio/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/", handler.Test)
	e.POST("/signUp", handler.SignUp)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return e
}
