package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type WinnersModel struct {
	WinnersModelId primitive.ObjectID `bson:"_id"`
	First          AWinnerModel       `json:"first,omitempty"`
	Second         AWinnerModel       `json:"second,omitempty"`
	Third          AWinnerModel       `json:"third,omitempty"`
}

type AWinnerModel struct {
	Uid            string `bson:"uid" json:"uid,omitempty"`
	RaffleNickName string
}
