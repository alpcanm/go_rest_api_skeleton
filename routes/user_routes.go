package routes

import (
	"go_rest_api_skeleton/controllers"
	"go_rest_api_skeleton/midlewares"

	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo) {

	e.POST("/users", controllers.InsertAUser)
	e.GET("/users", controllers.SelectAUser, midlewares.JwtSign())
}
