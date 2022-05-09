package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	Reply struct {
		ID        primitive.ObjectID `json:"_id" bson:"_id"`
		AdminId   primitive.ObjectID `json:"adminId" bson:"adminId"`
		ReviewId  primitive.ObjectID `json:"reviewId" bson:"reviewId"`
		Content   string             `json:"content" bson:"content"`
		CreatedAt string             `json:"createdAt" bson:"createdAt"`
		UpdatedAt string             `json:"updatedAt" bson:"updatedAt"`
	}
)
