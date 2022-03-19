package rafcont

import (
	"context"
	"fmt"
	"go_rest_api_skeleton/config"
	"go_rest_api_skeleton/models"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var rafflesCollection *mongo.Collection = config.GetCollection(config.DB, "raffles")
var usersCollection *mongo.Collection = config.GetCollection(config.DB, "users")
var rafflesWithSubscriberCollection *mongo.Collection = config.GetCollection(config.DB, "raffles-with-subscriber")
var validate = validator.New()

func InsertARaffle(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var raffle models.RaffleModel

	if err := c.Bind(&raffle); err != nil {
		//bir hata varsa döner
		fmt.Println(err.Error())
		return c.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
	}
	if err := validate.Struct(raffle); err != nil {
		fmt.Println(err.Error())

		return c.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
	}
	raffle.RaffleId = primitive.NewObjectID()
	raffle.IsExpired = false
	result, err := rafflesCollection.InsertOne(ctx, raffle)
	_, err2 := rafflesWithSubscriberCollection.InsertOne(ctx, raffle)

	if err != nil || err2 != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
	}
	//sonuç döner
	return c.JSON(http.StatusCreated, models.Response{Body: &echo.Map{"data": result}})

}
