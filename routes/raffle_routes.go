package routes

import (
	rafcont "go_rest_api_skeleton/controllers/raffle_contoller"

	"github.com/labstack/echo/v4"
)

func RaffleRoutes(e *echo.Echo) {
	e.POST("/raffles", rafcont.InsertARaffle)
	e.GET("/raffles", rafcont.GetRaffles)
	e.POST("/raffles-addto", rafcont.RaffleAddToListSubscribe)

}
