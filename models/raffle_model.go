package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type RaffleModel struct {
	RaffleId  primitive.ObjectID `bson:"_id" json:"raffle_id,omitempty"`
	Title     string             `json:"title,omitempty" validate:"required"`
	Comment   string             `json:"comment,omitempty" validate:"required"`
	Date      int64              `bson:"date" json:"date,omitempty" validate:"required"`
	PhotoUrl  string             `bson:"photo_url" json:"photo_url,omitempty" validate:"required"`
	IsExpired bool               `bson:"is_expired" json:"is_expired,omitempty"`
	Tag       string             `json:"tag,omitempty" validate:"required"`
	Url       string             `bson:"url" json:"url,omitempty" validate:"required"`
}

type MiniRaffleModel struct {
	MiniRaffleModelId primitive.ObjectID `bson:"_id"`
	RaffleId          primitive.ObjectID `bson:"raffle_id" json:"raffle_id,omitempty"`
	SubscribeId       primitive.ObjectID `bson:"subscribe_id"`
	SubscriberName    string             `json:"subscriber_name,omitempty" validate:"required"`
	SubscribeDate     int64              `json:"subscribe_date,omitempty" validate:"required"`
}
type RecentRaffleModel struct {
	RaffleId  primitive.ObjectID `bson:"raffle_id" json:"raffle_id,omitempty"`
	Title     string             `json:"title,omitempty" validate:"required"`
	Comment   string             `json:"comment,omitempty" validate:"required"`
	Date      int64              `bson:"date" json:"date,omitempty" validate:"required"`
	PhotoUrl  string             `bson:"photo_url" json:"photo_url,omitempty" validate:"required"`
	IsExpired bool               `bson:"is_expired" json:"is_expired,omitempty"`
	Tag       string             `json:"tag,omitempty" validate:"required"`
	Url       string             `bson:"url" json:"url,omitempty" validate:"required"`
}
