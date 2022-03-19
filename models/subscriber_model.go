package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type SubscriberModel struct {
	SubscribeId    primitive.ObjectID `bson:"_id"`
	SubscriberId   string             `bson:"subscriber_id" json:"subscriber_id,omitempty" validate:"required"`
	RaffleNickName string             `json:"raffle_nick_name,omitempty" validate:"required"`
	SubscribeDate  int64              `json:"subscribe_date,omitempty" validate:"required"`
}
