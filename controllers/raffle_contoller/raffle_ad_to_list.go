package rafcont

import (
	"context"
	"fmt"
	"go_rest_api_skeleton/models"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RaffleAddToListSubscribe(c echo.Context) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//! gelen istekten rfid parametresi alınıyor
	raffleId, _ := primitive.ObjectIDFromHex(c.QueryParam("rfid"))
	//!subscriberModel tanımlanır
	var subscriberModel models.SubscriberModel
	//!gelen veriler subscriber modele bind edilir.
	if err := c.Bind(&subscriberModel); err != nil {
		//!bir hata varsa döner
		fmt.Println(err.Error())
		return c.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
	}
	//! gelen veriler eksikliği kontrol ediliyor.
	if err := validate.Struct(subscriberModel); err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
	}
	//! subscriber a yeni ıd veriliyor
	subscriberModel.SubscribeModelId = primitive.NewObjectID()

	{ //! bu scope subscriber ı raffle listesi içerisine atıyor.

		filter := bson.D{{Key: "_id", Value: raffleId}}
		update := bson.D{{Key: "$push", Value: bson.D{{Key: "subscriber_list", Value: subscriberModel}}}}
		_, err := rafflesWithSubscriberCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
		}
	}

	{ //! bu scope mini raffle ı user ın listesi içerisine atıyor
		miniRaffle := models.MiniRaffleModel{
			MiniRaffleModelId: primitive.NewObjectID(),
			RaffleId:          raffleId,
			SubscribeId:       subscriberModel.SubscribeModelId,
			RaffleNickName:    subscriberModel.RaffleNickName,
			SubscribeDate:     subscriberModel.SubscribeDate,
		}

		filter := bson.D{{Key: "uid", Value: c.QueryParam("uid")}}
		update := bson.D{{Key: "$push", Value: bson.D{{Key: "subscribed_raffles", Value: miniRaffle}}}}
		_, err := usersCollection.UpdateOne(ctx, filter, update)
		if err != nil {

			return c.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
		}
	}

	return c.JSON(http.StatusOK, models.Response{Message: "", Body: &echo.Map{"data": "OK"}})
}
