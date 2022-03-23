package rafcont

import (
	"context"
	"go_rest_api_skeleton/models"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetARAffleFromRaffles(c echo.Context) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var raffle models.RaffleModel
	objectId, err1 := primitive.ObjectIDFromHex(c.QueryParam("raffleId"))
	if err1 != nil {
		if err1 == mongo.ErrNoDocuments {
			return c.JSON(http.StatusBadRequest, models.Response{Message: err1.Error()})
		}
	}

	err := rafflesCollection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&raffle)

	if err != nil {
		if err == mongo.ErrNoDocuments {

			return c.JSON(http.StatusNoContent, models.Response{Message: err.Error()})
		}

	}

	return c.JSON(http.StatusOK, models.Response{Body: &echo.Map{"data": raffle}})
}

func GetAWithWinners(c echo.Context) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var raffle bson.M
	objectId, err1 := primitive.ObjectIDFromHex(c.QueryParam("raffleId"))
	if err1 != nil {
		if err1 == mongo.ErrNoDocuments {
			return c.JSON(http.StatusBadRequest, models.Response{Message: err1.Error()})
		}
	}

	err := rafflesCollection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&raffle)

	if err != nil {
		if err == mongo.ErrNoDocuments {

			return c.JSON(http.StatusNoContent, models.Response{Message: err.Error()})
		}

	}

	return c.JSON(http.StatusOK, models.Response{Body: &echo.Map{"data": raffle}})
}
