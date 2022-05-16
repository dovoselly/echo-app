package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Reply struct {
		ID        primitive.ObjectID `json:"_id" bson:"_id"`
		UserId    primitive.ObjectID `json:"userId" bson:"userId"`
		ReviewId  primitive.ObjectID `json:"reviewId" bson:"reviewId"`
		Content   string             `json:"content" bson:"content"`
		CreatedAt string             `json:"createdAt" bson:"createdAt"`
		UpdatedAt string             `json:"updatedAt" bson:"updatedAt"`
	}

	CreateReply struct {
		Content string `json:"content"`
	}
)

func (r CreateReply) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Content, validation.Length(1, 5000).Error("length 1-5000 characters")),
	)
}
