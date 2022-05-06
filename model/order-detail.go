package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	OrderItem struct {
		ID        primitive.ObjectID `json:"_id" bson:"_id"`
		ProductId primitive.ObjectID `json:"productId" bson:"productId"`
		Quantity  int16              `json:"quantity" bson:"quantity"`
		Note      string             `json:"note" bson:"note"`
		Price     int64              `json:"price" bson:"price"`
		CreatedAt string             `json:"createdAt" bson:"createdAt"`
		UpdatedAt string             `json:"updatedAt" bson:"updatedAt"`
	}
)
