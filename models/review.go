package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Review struct {
		ID        primitive.ObjectID `json:"_id" bson:"_id"`
		UserId    primitive.ObjectID `json:"userId" bson:"userId"`
		ProductId primitive.ObjectID `json:"productId" bson:"productId"`
		Rating    uint               `json:"rating" bson:"rating"`
		Content   string             `json:"content" bson:"content"`
		CreatedAt string             `json:"createdAt" bson:"createdAt"`
		UpdatedAt string             `json:"updatedAt" bson:"updatedAt"`
	}

	QueryReview struct {
		Page      string `query:"page"`
		Rating    uint   `query:"rating"`
		CreatedAt string `query:"createdAt"`
	}

	CreateReview struct {
		Rating  string `json:"rating"`
		Content string `json:"content"`
	}
)

func (r CreateReview) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Rating, validation.Length(1, ))
	)
}
