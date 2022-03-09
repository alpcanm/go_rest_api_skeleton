package models

import (
	"github.com/labstack/echo/v4"
)

type Response struct {
	Message string    `json:"message"`
	Body    *echo.Map `json:"body"`
}
