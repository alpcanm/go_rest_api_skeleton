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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetSomeSubscribersFromRecentRaffle(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	recentRaffle := getRecentRaffle()
	var resultRaffle struct { //sadece içerisindeki listeyi çeker
		SubscriberList []models.SubscriberModel `bson:"subscriber_list" json:"subscriber_list"`
	}
	winnersChoose := checkWinnersChooss(c.QueryParam("winnersChoose"))
	generateNumber := checkNumber(winnersChoose, c)
	err := rafflesWithSubscriberCollection.FindOne(ctx, bson.M{"_id": recentRaffle.RaffleId}).Decode(&resultRaffle)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.JSON(http.StatusNoContent, models.Response{Message: err.Error()})
		}

	}
	subscriberList := resultRaffle.SubscriberList

	indexesOfChoices := randomGenerator(len(subscriberList), generateNumber)

	choosens := chooseSomeSubscribers(indexesOfChoices, subscriberList)
	if winnersChoose {
		//Eğer winnerChoose true ise kazananları raffleCollections içerisindeki o raffle içerisine winnerModel olarak ekleyecek.
		setWinners(choosens, recentRaffle.RaffleId)
	}
	return c.JSON(http.StatusOK, models.Response{Body: &echo.Map{"data": choosens}})

}
func checkNumber(winnerChoose bool, c echo.Context) int {
	if winnerChoose {
		return 3
	}
	yNumberValue, err := strconv.Atoi(c.QueryParam("number"))
	if err != nil {
		panic(err)
	}
	return yNumberValue
}
func checkWinnersChooss(param string) bool {
	if param == "true" {
		return true
	}
	return false
}
func chooseSomeSubscribers(indexList []int, subscrierList []models.SubscriberModel) []models.SubscriberModel {
	var resultList []models.SubscriberModel
	for _, num := range indexList {
		resultList = append(resultList, subscrierList[num])
	}
	return resultList
}

func randomGenerator(x int, y int) []int {
	// 0 ile x arasında y adet birbirinden farklı sayıları liste içinde döndüren fonksisyon
	rand.Seed(time.Now().UnixNano())
	var generatedNumbers []int

	isNumberIn := func(list []int, sayi int) bool {
		for _, num := range list {
			if num == sayi {

				return true
			}
		}
		return false
	}

	i := 0
	for i < y {
		randomNumber := rand.Intn(x)
		if !isNumberIn(generatedNumbers, randomNumber) {
			i++
			generatedNumbers = append(generatedNumbers, randomNumber)
		}
	}
	return generatedNumbers
}

func setWinners(winners []models.SubscriberModel, objectId primitive.ObjectID) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var winnerModel models.WinnersModel
	winnerModel.WinnersModelId = primitive.NewObjectID()
	winnerModel.First = winners[0]
	winnerModel.Second = winners[1]
	winnerModel.Third = winners[2]

	filter := bson.D{{Key: "_id", Value: objectId}}
	update := bson.D{{Key: "$push", Value: bson.D{{Key: "winners", Value: winnerModel}}}}
	_, err := rafflesCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		panic(err)
	}
	return true

}
