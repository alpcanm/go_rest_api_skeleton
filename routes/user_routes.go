package routes

import (
	"go_rest_api_skeleton/controllers"

	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo) {
	e.POST("/user", controllers.InsertAUser)
	e.GET("/user/:uid", controllers.SelectAUser)
}
