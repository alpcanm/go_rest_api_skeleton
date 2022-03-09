package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductModel struct {
	ProductId primitive.ObjectID `bson:"product_id" json:"product_id,omitempty"`
	Title     string             `json:"title,omitempty" validate:"required"`
	Comment   string             `json:"comment,omitempty" validate:"required"`
	DrawDate  int64              `bson:"draw_date" json:"draw_date,omitempty" validate:"required"`
	PhotoUrl  string             `bson:"photo_url" json:"photo_url,omitempty" validate:"required"`
	IsExpired bool               `bson:"is_expired" json:"is_expired,omitempty"`
	Tag       string             `json:"tag,omitempty" validate:"required"`
	DrawUrl   string             `bson:"draw_url" json:"draw_url,omitempty" validate:"required"`
}

type ProductFilterModel struct {
	ProductId string             `bson:"product_id" json:"product_id"`
	DrawDate  primitive.DateTime `bson:"draw_date" json:"draw_date"`
	Tag       string             `json:"tag"`
	IsExpired string             `bson:"is_expired" json:"is_expired"`
}
