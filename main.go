package main

import (
	"go_rest_api_skeleton/config"
	"go_rest_api_skeleton/midlewares"
	"go_rest_api_skeleton/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	midlewares.MainMiddlewares(e)
	routes.ApiKeyRoute(e)
	routes.UserRoute(e)
	routes.RaffleRoutes(e)
	e.Logger.Fatal(e.Start(config.LocalHost()))
}
