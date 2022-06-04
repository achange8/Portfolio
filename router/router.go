package router

import (
	"net/http"

	"github.com/achange8/Portfolio/handler"
	"github.com/achange8/Portfolio/middlewares"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//board crud done
//user curd	done
//oauth google login not done
func New() *echo.Echo {
	e := echo.New()
	g := e.Group("/board")
	g.Use(middlewares.TokenCheckMiddleware)

	e.GET("/", handler.Test)
	e.POST("/signUp", handler.SignUp)         //done
	e.POST("/signIn", handler.SignIn)         //done
	e.POST("/modifyID", handler.ModifyID)     //done
	e.POST("/modifyPW", handler.ModifyPW)     //done
	e.POST("/duplicate", handler.DuplCheckID) //done
	e.GET("/signOut", handler.SignOut)        //done
	g.POST("/write", handler.CreateBoard)     //done
	e.GET("/listBoard", handler.ListBoard)    //done
	e.GET("/readBoard/", handler.ReadBoard)   //done
	g.POST("/modify/", handler.UpdateBoard)   //done
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	})) //CORS setting
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.DELETE("/user", handler.UserDelete) //testing
	e.GET("/auth/google/login", handler.GoogleLogin)
	e.GET("/auth/google/callback", handler.GoogleCallBack)

	return e
}
