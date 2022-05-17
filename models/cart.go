package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	CartBSON struct {
		ID        primitive.ObjectID `bson:"_id"`
		UserId    primitive.ObjectID `bson:"userId"`
		Items     []CartItemBSON     `bson:"items"`
		CreatedAt time.Time          `bson:"createdAt"`
		UpdatedAt time.Time          `bson:"updatedAt"`
	}

	CartCreateBSON struct {
		ID        primitive.ObjectID   `bson:"_id"`
		UserId    primitive.ObjectID   `bson:"userId"`
		Items     []primitive.ObjectID `bson:"items"`
		CreatedAt time.Time            `bson:"createdAt"`
		UpdatedAt time.Time            `bson:"updatedAt"`
	}

	CartCreate struct {
		UserId primitive.ObjectID `bson:"userId"`
		Items  []CartItemsCreate  `bson:"items"`
	}
	CartResponse struct {
		ID        primitive.ObjectID `bson:"_id"`
		UserId    primitive.ObjectID `bson:"userId"`
		Items     []CartItemResponse `bson:"items"`
		CreatedAt time.Time          `bson:"createdAt"`
		UpdatedAt time.Time          `bson:"updatedAt"`
	}

	CartUpdate struct {
		Quantity int16 `bson:"quantity"`
	}
)
