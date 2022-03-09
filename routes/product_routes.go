package routes

import (
	prodcont "go_rest_api_skeleton/controllers/product_contoller"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Echo) {
	e.POST("/products", prodcont.InsertAProduct)
	e.GET("/products", prodcont.GetProducts)
}
