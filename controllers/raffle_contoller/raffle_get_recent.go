package rafcont

import (
	"context"
	"go_rest_api_skeleton/models"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetRecentRaffle(c echo.Context) error {

	return c.JSON(http.StatusOK, models.Response{Body: &echo.Map{"data": getRecentRaffle()}})
}

func getRecentRaffle() models.RecentRaffleModel {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var raffle models.RecentRaffleModel

	err := recentRaffleColl.FindOne(ctx, bson.M{"is_expired": false}).Decode(&raffle)
	if err != nil {
		if err == mongo.ErrNoDocuments {

			panic(err)
		}

	}
	return raffle
}
