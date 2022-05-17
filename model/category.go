package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	CategoryBSON struct {
		ID          primitive.ObjectID `bson:"_id"`
		Name        string             `bson:"name"`
		Description string             `bson:"description"`
		Status      string             `bson:"status"`
		CreatedAt   time.Time          `bson:"createdAt"`
		UpdatedAt   time.Time          `bson:"updatedAt"`
	}
	CategoryResponse struct {
		ID          primitive.ObjectID `json:"_id"`
		Name        string             `json:"name"`
		Description string             `json:"description"`
		Status      string             `json:"status"`
		CreatedAt   time.Time          `json:"createdAt"`
		UpdatedAt   time.Time          `json:"updatedAt"`
	}

	CategoryCreateBody struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Status      string `json:"status"`
	}

	CategoryUpdateBody struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
)

func (c CategoryCreateBody) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Required),
		validation.Field(&c.Description, validation.Required),
	)
}

func (c CategoryUpdateBody) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Required),
		validation.Field(&c.Description, validation.Required),
	)
}

// NewCategory ...
//func (c CategoryCreateBody) NewCategory() Category {
//	return Category{
//		ID         : primitive.NewObjectID(),
//		Name      : c.Name,
//		Description string             `bson:"description"`
//
//		bool:
//			boo
//		Status     :  ,
//		CreatedAt   :
//		UpdatedAt   time.Time          `bson:"updatedAt"`
//	}
//}
