package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type SubscriberModel struct {
	SubscribeModelId primitive.ObjectID `bson:"_id"`
	SubscriberId     string             `bson:"subscriber_id" json:"subscriber_id,omitempty" validate:"required"`
	RaffleNickName   string             `bson:"raffle_nick_name" json:"raffle_nick_name,omitempty" validate:"required"`
	SubscribeDate    int64              `bson:"subscribe_date"  json:"subscribe_date,omitempty" validate:"required"`
}

type WithIndexSubscriberModel struct {
	SubscribeModelId primitive.ObjectID `json:"_id"`
	SubscriberId     string             `json:"subscriber_id,omitempty"`
	RaffleNickName   string             `json:"raffle_nick_name,omitempty" `
	SubscribeDate    int64              `json:"subscribe_date"`
	Index            int                `json:"index"`
}
