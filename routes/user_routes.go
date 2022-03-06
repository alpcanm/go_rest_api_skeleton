package routes

import (
	"go_rest_api_skeleton/controllers"

	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo) {
	e.POST("/user", controllers.InsertOneUser)
	e.GET("/user/:userId", controllers.GetAUser)
	e.PUT("/user/:userId", controllers.UpdateAUser)
}
