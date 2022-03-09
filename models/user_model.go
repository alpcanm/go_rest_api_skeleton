package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserModel struct {
	Id           primitive.ObjectID `json:"id,omitempty"`
	Name         string             `json:"name,omitempty" validate:"required"`
	Surname      string             `json:"surname,omitempty" validate:"required"`
	Mail         string             `json:"mail,omitempty" validate:"required"`
	Uid          string             `json:"uid,omitempty" validate:"required"`
	PhoneNumber  string             `bson:"phone_number" json:"phone_number,omitempty"`
	MailVerified bool               `bson:"mail_verified" json:"mail_verified,omitempty"`
	DrawNickName string             `bson:"draw_nick_name" json:"draw_nick_name,omitempty"`
	Tickets      int                `json:"tickets,omitempty"`
	CreatedAt    primitive.DateTime `bson:"created_at" json:"created_at,omitempty"`
}
