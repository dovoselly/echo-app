package models

import "go.mongodb.org/mongo-driver/bson/primitive"

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
		Page      string `json:"page" bson:"page"`
		Rating    uint   `json:"rating" bson:"rating"`
		CreatedAt string `json:"createdAt" bson:"createdAt"`
	}
)
