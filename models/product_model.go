package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ProductId primitive.ObjectID `json:"product_id,omitempty"`
	Title     string             `json:"title,omitempty" validate:"required"`
	Comment   string             `json:"comment,omitempty" validate:"required"`
	DrawDate  primitive.DateTime `json:"draw_date,omitempty" validate:"required"`
	PhotoUrl  string             `json:"photo_url,omitempty" validate:"required"`
	IsExpired bool               `json:"is_expired,omitempty" validate:"required"`
	Tax       string             `json:"tag,omitempty" validate:"required"`
	DrawUrl   int                `json:"draw_url,omitempty" validate:"required"`
}
