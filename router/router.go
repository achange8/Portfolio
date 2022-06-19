package router

import (
	"net/http"

	"github.com/achange8/Portfolio/handler"
	"github.com/achange8/Portfolio/middlewares"
	"github.com/achange8/Portfolio/module"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//board crud done
//user curd	done
//oauth google login done
//todos : search api
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
	g.DELETE("/delete/", handler.DeleteBoard) // done
	e.DELETE("/user", handler.UserDelete)     //done
	///for test user info///
	e.GET("/allUser", handler.GetAllUsers) //done
	e.GET("/search", handler.SearchBoard)  // done
	////

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	})) //CORS setting
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//Google Oauth login
	e.GET("/auth/google/login", handler.GoogleLogin)       //done
	e.GET("/auth/google/callback", handler.GoogleCallBack) //done

	e.POST("/upload", module.Upload)          // done
	e.GET("/download/", handler.DownLoadFile) // testing
	e.GET("/load/", handler.LoadFile)

	return e
}
