package midlewares

import (
	"go_rest_api_skeleton/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JwtSign() echo.MiddlewareFunc {
	return middleware.JWT([]byte(config.JwtKey()))
}
