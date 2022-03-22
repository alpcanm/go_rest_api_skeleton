package routes

import (
	rafcont "go_rest_api_skeleton/controllers/raffle_contoller"

	"github.com/labstack/echo/v4"
)

func RaffleRoutes(e *echo.Echo) {
	e.POST("/raffle-recent", rafcont.SetNewRecentRaffle)
	e.PATCH("/raffle-recent", rafcont.RaffleRecentSetExpired)
	e.GET("/raffle-recent", rafcont.GetRecentRaffle)

	e.GET("/raffle", rafcont.GetARAffleFromRaffles)
	e.POST("/raffles", rafcont.InsertARaffle)
	e.GET("/raffles", rafcont.GetAllRaffles)

	e.POST("/raffles-addto", rafcont.RaffleAddToListSubscribe)

	e.GET("/recent-raffle-some-subscribers", rafcont.GetSomeSubscribersFromRecentRaffle)

}
