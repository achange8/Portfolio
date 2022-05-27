package router

import (
	"net/http"

	"github.com/achange8/Portfolio/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/", handler.Test)
	e.POST("/signUp", handler.SignUp)         //done
	e.POST("/signIn", handler.SignIn)         //done
	e.POST("/modifyID", handler.ModifyID)     //done
	e.POST("/modifyPW", handler.ModifyPW)     //done
	e.POST("/duplicate", handler.DuplCheckID) //done
	e.POST("/write", handler.CreateBoard)     //testing
	e.GET("/signOut", handler.SignOut)        //done

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	})) //CORS setting
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return e
}
