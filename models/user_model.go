package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Uid          primitive.ObjectID `json:"uid,omitempty"`
	Name         string             `json:"name,omitempty" validate:"required"`
	Surname      string             `json:"surname,omitempty" validate:"required"`
	Mail         string             `json:"mail,omitempty" validate:"required"`
	PhoneNumber  string             `json:"phone_number,omitempty"`
	MailVerified bool               `json:"mail_verified,omitempty" validate:"required"`
	DrawNickName string             `json:"draw_nick_name,omitempty" validate:"required"`
	Tickets      int                `json:"tickets,omitempty" validate:"required"`
	CreatedAt    primitive.DateTime `json:"created_at,omitempty" validate:"required"`
}
