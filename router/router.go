package router

import (
	"net/http"

	"github.com/achange8/Portfolio/handler"
	"github.com/achange8/Portfolio/handler/board"
	"github.com/achange8/Portfolio/handler/user"
	"github.com/achange8/Portfolio/middlewares"
	"github.com/achange8/Portfolio/module"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//board crud done
//user curd	done
//oauth google login done
//search api done

//todos : db refactorying , transaction user,board
func New() *echo.Echo {
	e := echo.New()
	g := e.Group("/board")
	u := e.Group("/user")

	u.Use(middlewares.TokenCheckMiddleware)
	g.Use(middlewares.TokenCheckMiddleware)

	u.GET("/check", user.Usercheck)
	e.GET("/", handler.Test)
	//user sign
	e.POST("/signUp", user.SignUp)         //done
	e.POST("/signIn", user.SignIn)         //done
	e.POST("/modifyID", user.ModifyID)     //done
	e.POST("/modifyPW", user.ModifyPW)     //done
	e.POST("/duplicate", user.DuplCheckID) //done
	e.GET("/signOut", user.SignOut)        //done
	e.DELETE("/user", user.UserDelete)     //done
	//board
	e.GET("/search", board.SearchBoard)     // done
	e.GET("/list", board.ListBoard)         //done
	e.GET("/readBoard/", board.ReadBoard)   //done
	g.POST("/write", board.CreateBoard)     //done
	g.POST("/modify/", board.UpdateBoard)   //done
	g.DELETE("/delete/", board.DeleteBoard) // done
	///for test user info///
	e.GET("/allUser", user.GetAllUsers) //done
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

	//file upload
	e.POST("/upload", module.Upload)        // done
	e.GET("/download/", board.DownLoadFile) // done
	e.GET("/load/", board.LoadFile)

	return e
}
