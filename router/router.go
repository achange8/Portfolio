package router

import (
	"github.com/achange8/Portfolio/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/", handler.Test)
	e.POST("/signUp", handler.SignUp)         //done
	e.POST("/signIn", handler.SignIn)         // done
	e.POST("/modifyID", handler.ModifyID)     //done
	e.POST("/modifyPW", handler.ModifyPW)     //testing
	e.POST("/duplicate", handler.DuplCheckID) //done
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return e
}
