package routes

import (
	"go_rest_api_skeleton/controllers"

	"github.com/labstack/echo/v4"
)

func ApiKeyRoute(e *echo.Echo) {
	e.GET("/api-key", controllers.GetApiKey)
}
