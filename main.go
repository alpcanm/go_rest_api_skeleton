package main

import (
	config_ "go_rest_api_skeleton/config"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	config_.InitializeSetup()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start("127.0.0.1:3000"))
}
