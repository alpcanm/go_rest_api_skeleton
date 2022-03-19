package routes

import (
	rafcont "go_rest_api_skeleton/controllers/raffle_contoller"

	"github.com/labstack/echo/v4"
)

func RaffleRoutes(e *echo.Echo) {
	e.GET("/raffle", rafcont.GetARAffle)
	e.POST("/raffle", rafcont.SetNewRecentRaffle)
	e.GET("/raffle-recent", rafcont.GetRecentRaffle)
	e.POST("/raffles", rafcont.InsertARaffle)
	e.GET("/raffles", rafcont.GetAllRaffles)
	e.POST("/raffles-addto", rafcont.RaffleAddToListSubscribe)

}
