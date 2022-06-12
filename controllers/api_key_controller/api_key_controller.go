package apicont

import (
	"go_rest_api_skeleton/config"
	"go_rest_api_skeleton/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetApiKey(c echo.Context) error {
	//! firebase api key ini getiren cevap
	return c.JSON(http.StatusOK, models.Response{Body: &echo.Map{"data": config.ApiKey()}})
}
