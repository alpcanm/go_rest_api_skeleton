package rafcont

import (
	"context"
	"go_rest_api_skeleton/models"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetSomeSubscribersFromRecentRaffle(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	recentRaffle := getRecentRaffle()
	var resultRaffle struct { //sadece içerisindeki listeyi çeker
		SubscriberList []models.SubscriberModel `bson:"subscriber_list" json:"subscriber_list"`
	}

	err := rafflesWithSubscriberCollection.FindOne(ctx, bson.M{"_id": recentRaffle.RaffleId}).Decode(&resultRaffle)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.JSON(http.StatusNoContent, models.Response{Message: err.Error()})
		}

	}
	subscriberList := resultRaffle.SubscriberList

	indexesOfChoices := randomGenerator(len(subscriberList), c.QueryParam("number"))
	choosens := chooseSomeSubscribers(indexesOfChoices, subscriberList)
	return c.JSON(http.StatusOK, models.Response{Body: &echo.Map{"data": choosens}})

}

func chooseSomeSubscribers(indexList []int, subscrierList []models.SubscriberModel) []models.SubscriberModel {
	var resultList []models.SubscriberModel
	for _, num := range indexList {
		resultList = append(resultList, subscrierList[num])
	}
	return resultList
}

func randomGenerator(x int, y string) []int {
	// 0 ile x arasında y adet birbirinden farklı sayıları liste içinde döndüren fonksisyon
	rand.Seed(time.Now().UnixNano())
	var generatedNumbers []int
	yNumberValue, err := strconv.Atoi(y)
	if err != nil {
		panic(err)
	}
	isNumberIn := func(list []int, sayi int) bool {
		for _, num := range list {
			if num == sayi {

				return true
			}
		}
		return false
	}

	i := 0
	for i < yNumberValue {
		randomNumber := rand.Intn(x)
		if !isNumberIn(generatedNumbers, randomNumber) {
			i++
			generatedNumbers = append(generatedNumbers, randomNumber)
		}
	}
	return generatedNumbers
}
