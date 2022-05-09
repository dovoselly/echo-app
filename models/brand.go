package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	BrandBSON struct {
		ID          primitive.ObjectID `bson:"_id"`
		Name        string             `bson:"name"`
		Description string             `bson:"description"`
		Status      string             `bson:"status"`
		CreatedAt   time.Time          `bson:"createdAt"`
		UpdatedAt   time.Time          `bson:"updatedAt"`
	}

	BrandResponse struct {
		ID          primitive.ObjectID `json:"_id"`
		Name        string             `json:"name"`
		Description string             `json:"description"`
		Status      string             `json:"status"`
		CreatedAt   time.Time          `json:"createdAt"`
		UpdatedAt   time.Time          `json:"updatedAt"`
	}

	BrandCreate struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	BrandUpdate struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
)
