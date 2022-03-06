package main

import (
	"go_rest_api_skeleton/routes"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	routes.UserRoute(e)
	e.Logger.Fatal(e.Start("127.0.0.1:3000"))
}
