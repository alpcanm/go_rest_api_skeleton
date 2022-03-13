package routes

import (
	usercont "go_rest_api_skeleton/controllers/user_controller"
	"go_rest_api_skeleton/midlewares"

	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo) {

	e.POST("/users", usercont.InsertAUser)
	e.GET("/users", usercont.SelectAUser, midlewares.JwtSign())
	e.GET("/users/:uid", usercont.SelectUsersRaffles)
}
