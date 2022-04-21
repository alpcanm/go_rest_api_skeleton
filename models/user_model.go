package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserModel struct {
	Id                primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Name              string             `json:"name,omitempty" validate:"required"`
	Mail              string             `json:"mail,omitempty" validate:"required"`
	Uid               string             `json:"uid,omitempty" validate:"required"`
	MailVerified      bool               `bson:"mail_verified" json:"mail_verified,omitempty"`
	RaffleNickName    string             `bson:"raffle_nick_name" json:"raffle_nick_name,omitempty"`
	CreatedAt         int64              `bson:"created_at" json:"created_at,omitempty"`
	SubscribedRaffles []MiniRaffleModel  `bson:"subscribed_raffles" json:"subscribed_raffles,omitempty"`
}

type UsersRaffleList struct {
	RaffleList []RaffleModel `json:"subscribed_raffles,omitempty"`
}
