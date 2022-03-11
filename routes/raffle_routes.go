package routes

import (
	prodcont "go_rest_api_skeleton/controllers/raffle_contoller"

	"github.com/labstack/echo/v4"
)

func RaffleRoutes(e *echo.Echo) {
	e.POST("/raffles", prodcont.InsertARaffle)
	e.GET("/raffles", prodcont.GetRaffles)
	e.GET("/dene", prodcont.DenemeProd)
}
