package routes

import (
	apicont "go_rest_api_skeleton/controllers/api_key_controller"

	"github.com/labstack/echo/v4"
)

func ApiKeyRoute(e *echo.Echo) {
	e.GET("/api-key", apicont.GetApiKey)
}
