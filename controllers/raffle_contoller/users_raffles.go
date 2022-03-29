package rafcont

import (
	"context"
	"fmt"
	"go_rest_api_skeleton/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetUsersRaffles(greaterThan int, wantedRaffles []primitive.ObjectID) []models.RaffleModel {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var results []models.RaffleModel
	sortValue := bson.D{{Key: date, Value: 1}}
	opts := options.Find().SetSort(sortValue)

	fetchedData, err := rafflesCollection.Find(ctx, getFiltereds(greaterThan, wantedRaffles), opts)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("Boş veri")
		}
		panic(err)
	}

	for fetchedData.Next(ctx) {

		var singleRaffle models.RaffleModel
		if err = fetchedData.Decode(&singleRaffle); err != nil {
			panic(err)
		}

		results = append(results, singleRaffle)
	}
	return results
}

func getFiltereds(gtVal int, wantedRaffles []primitive.ObjectID) bson.D {

	greaterThanValue := bson.D{{Key: gt, Value: gtVal}}
	greaterThan := primitive.E{Key: date, Value: greaterThanValue}

	var equalList []bson.D

	for _, value := range wantedRaffles {
		equalList = append(equalList, bson.D{{Key: "_id", Value: bson.D{{Key: eq, Value: value}}}})
	}
	return bson.D{greaterThan, {Key: "$or", Value: equalList}}

}
