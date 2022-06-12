package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func connectDB() *mongo.Client {

	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	//! Veri tabanına bağlanır
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI()))

	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Bağlantı başarılı")

	return client
}

var DB *mongo.Client = connectDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("mesela-service").Collection(collectionName)
	return collection
}
