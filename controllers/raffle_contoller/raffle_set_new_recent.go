package rafcont

import (
	"context"
	"go_rest_api_skeleton/config"
	"go_rest_api_skeleton/models"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var recentRaffleColl *mongo.Collection = config.GetCollection(config.DB, "recent-raffle")

//! recent raffle ı atadığımız istek.
func SetNewRecentRaffle(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	opt := options.Find().SetSort(bson.M{"date": 1}).SetLimit(1)                      //!en küçük olanı getirir.Yani tarihi en yakın olan. Çünkü tarihi milisecond cinsinden giriyoruz.
	fetchedData, err := rafflesCollection.Find(ctx, bson.M{"is_expired": false}, opt) //! Aynı zamanda çekilişn is_expired false olmalı. Bu çekilişin henüz yapılmadığını gösterir.

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
		}
	}
	var results []models.RaffleModel
	for fetchedData.Next(ctx) {
		var singleRaffle models.RaffleModel
		// fmt.Println(singleRaffle)
		if err = fetchedData.Decode(&singleRaffle); err != nil {
			return c.JSON(http.StatusInternalServerError, models.Response{Message: "error", Body: &echo.Map{"data": err.Error()}})
		}
		results = append(results, singleRaffle)
	}
	//! çoklu sorgu yapmışız gibi bir liste döndürüyor ancak bu liste tek elemanlı onun sıfırıncı elemanı bizim yeni recentRaffle ımız.
	fetchedRaffle := results[0]
	//! fetchedRaffle ı recentRaffle isminde yeni bir instance a dönüştürüyoruz.
	recentRaffle := models.RecentRaffleModel{
		RaffleId:  fetchedRaffle.RaffleId,
		Title:     fetchedRaffle.Title,
		Comment:   fetchedRaffle.Comment,
		Date:      fetchedRaffle.Date,
		PhotoUrl:  fetchedRaffle.PhotoUrl,
		IsExpired: fetchedRaffle.IsExpired,
		Tag:       fetchedRaffle.Tag,
		Url:       fetchedRaffle.Url,
	}
	filter := bson.M{"is_expired": true}
	//! recentRaffle instance ımızı recentRaffle collections ındaki veriyle değiştiriyoruz. ReplaceOne metoduyla.
	result, err := recentRaffleColl.ReplaceOne(ctx, filter, recentRaffle)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Body: &echo.Map{"data": err.Error()}})
	}
	return c.JSON(http.StatusOK, models.Response{Body: &echo.Map{"data": result}})
}
