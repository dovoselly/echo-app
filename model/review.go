package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	Review struct {
		ID        primitive.ObjectID `json:"_id" bson:"_id"`
		AdminId   primitive.ObjectID `json:"adminId" bson:"adminId"`
		ReviewId  primitive.ObjectID `json:"reviewId" bson:"reviewId"`
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
