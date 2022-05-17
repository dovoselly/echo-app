package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Review struct {
		ID        primitive.ObjectID `json:"_id" bson:"_id"`
		UserId    primitive.ObjectID `json:"userId" bson:"userId"`
		ProductId primitive.ObjectID `json:"productId" bson:"productId"`
		Rating    uint8              `json:"rating" bson:"rating"`
		Content   string             `json:"content" bson:"content"`
		CreatedAt string             `json:"createdAt" bson:"createdAt"`
		UpdatedAt string             `json:"updatedAt" bson:"updatedAt"`
	}

	ReviewQuery struct {
		Page   int64  `query:"page"`
		Rating string `query:"rating"`
		Sort   string `query:"sort"`
	}

	CreateReview struct {
		Content string `json:"content"`
		Rating  uint8  `json:"rating"`
	}
)

func (r CreateReview) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Content, validation.Length(1, 5000).Error("length 1-5000 characters")),
	)
}
