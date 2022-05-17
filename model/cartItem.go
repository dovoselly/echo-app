package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	CartItemBSON struct {
		ID        primitive.ObjectID `bson:"_id"`
		ProductId primitive.ObjectID `bson:"productId"`
		Quantity  uint               `bson:"quantity"`
		Note      string             `bson:"note"`
		Price     string             `bson:"price"`
		CreatedAt time.Time          `bson:"createdAt"`
		UpdatedAt time.Time          `bson:"updatedAt"`
	}

	CartItemResponse struct {
		ID        primitive.ObjectID `json:"_id"`
		ProductId primitive.ObjectID `json:"productId"`
		Quantity  uint               `json:"quantity"`
		Note      string             `json:"note"`
		Price     string             `json:"price"`
		CreatedAt time.Time          `json:"createdAt"`
		UpdatedAt time.Time          `json:"updatedAt"`
	}

	CartItemsCreate struct {
		ProductId primitive.ObjectID `json:"productId"`
		Quantity  uint               `json:"quantity"`
		Note      string             `json:"note"`
		Price     string             `json:"price"`
	}
)
