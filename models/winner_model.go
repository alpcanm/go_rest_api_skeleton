package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type WinnersModel struct {
	WinnersModelId primitive.ObjectID `bson:"_id"`
	First          SubscriberModel    `json:"first,omitempty"`
	Second         SubscriberModel    `json:"second,omitempty"`
	Third          SubscriberModel    `json:"third,omitempty"`
}
