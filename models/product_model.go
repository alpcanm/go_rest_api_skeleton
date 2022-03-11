package models

type RaffleModel struct {
	Title     string `json:"title,omitempty" validate:"required"`
	Comment   string `json:"comment,omitempty" validate:"required"`
	Date      int64  `bson:"date" json:"date,omitempty" validate:"required"`
	PhotoUrl  string `bson:"photo_url" json:"photo_url,omitempty" validate:"required"`
	IsExpired bool   `bson:"is_expired" json:"is_expired,omitempty"`
	Tag       string `json:"tag,omitempty" validate:"required"`
	Url       string `bson:"url" json:"url,omitempty" validate:"required"`
}
