package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"

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

	BrandCreateBody struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	BrandUpdateBody struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
)

func (c BrandCreateBody) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Required),
		validation.Field(&c.Description, validation.Required),
	)
}

func (c BrandUpdateBody) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Required),
		validation.Field(&c.Description, validation.Required),
	)
}
