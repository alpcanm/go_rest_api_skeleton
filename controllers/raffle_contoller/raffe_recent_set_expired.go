package rafcont

import (
	"context"
	"go_rest_api_skeleton/models"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func RaffleRecentSetExpired(c echo.Context) error {
	//! recent raffle'i expired hale getiren istek. Çekiliş bittikten sonra çalışacak istek.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//! Veriye ulaşıp is_expired field ı false olanı buluyoruz. Aynı zamanda zamana göre sıraladığı için tarihi en kısa olanı seçiyor.
	filter := bson.M{"is_expired": false}
	//! is_expired false ve aynı zamanda
	update := bson.M{"$set": bson.M{"is_expired": true}}
	//!Aynı anda 3 collections da değiştirilmesi gerek.
	_, err1 := rafflesCollection.UpdateOne(ctx, filter, update)
	_, err2 := rafflesWithSubscriberCollection.UpdateOne(ctx, filter, update)
	result, err := recentRaffleColl.UpdateOne(ctx, filter, update)

	if err != nil || err1 != nil || err2 != nil {
		return c.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, models.Response{Body: &echo.Map{"": result}})
}
