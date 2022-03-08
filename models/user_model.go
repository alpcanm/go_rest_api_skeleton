package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserModel struct {
	Id           primitive.ObjectID `json:"id,omitempty"`
	Name         string             `json:"name,omitempty" validate:"required"`
	Surname      string             `json:"surname,omitempty" validate:"required"`
	Mail         string             `json:"mail,omitempty" validate:"required"`
	Uid          string             `json:"uid,omitempty" validate:"required"`
	PhoneNumber  string             `json:"phone_number,omitempty"`
	MailVerified bool               `json:"mail_verified,omitempty"`
	DrawNickName string             `json:"draw_nick_name,omitempty"`
	Tickets      int                `json:"tickets,omitempty"`
	CreatedAt    primitive.DateTime `json:"created_at,omitempty"`
}
