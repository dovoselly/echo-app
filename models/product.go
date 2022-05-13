package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Product struct {
		Id          primitive.ObjectID `json:"_id" bson:"_id"`
		Name        string             `json:"name" bson:"name"`
		CategoryId  primitive.ObjectID `json:"category" bson:"category"`
		BrandId     primitive.ObjectID `json:"brandId" bson:"brandId"`
		Price       uint               `json:"price" bson:"price"`
		Description string             `json:"description" bson:"description"`
		Images      []string           `json:"images" bson:"images"`
		Quantity    uint               `json:"quantity" bson:"quantity"`
		Rest        uint               `json:"rest" bson:"rest"`
		Status      string             `json:"status" bson:"status"`
		CreatedAt   string             `json:"createdAt" bson:"createdAt"`
		UpdatedAt   string             `json:"updatedAt" bson:"updatedAt"`
	}

	ProductQuery struct {
		Page        int64  `query:"page" binding:"required"`
		Name        string `query:"name,omitempty"`
		CategoryId  string `query:"categoryId,omitempty"`
		PriceFromTo string `query:"priceFromTo,omitempty"`
		BrandId     string `query:"brandId,omitempty"`
		Sort        string `query:"sort,omitempty"`
		Price       string `query:"price,omitempty"`
		CreatedAt   string `query:"createdAt,omitempty"`
	}
)

func (p Product) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Required),
		validation.Field(&p.Description, validation.Required),
		validation.Field(&p.Images, validation.Length(1, 20)),
	)
}
