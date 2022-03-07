package controllers

import (
	"fmt"
	"go_rest_api_skeleton/config"
	"go_rest_api_skeleton/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetApiKey(c echo.Context) error {
	fmt.Println("Api key istek")
	return c.JSON(http.StatusOK, models.Response{Body: &echo.Map{"data": config.ApiKey()}})
}
