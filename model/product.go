package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	// FIXME: Tách ra nhiều struct
	Product struct {
		Id          primitive.ObjectID `json:"_id" bson:"_id"`
		Name        string             `json:"name" bson:"name"`
		CategoryId  primitive.ObjectID `json:"categoryId" bson:"categoryId"`
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
		Page       int64  `query:"page" binding:"required"`
		Name       string `query:"name"`
		CategoryId string `query:"categoryId"`
		PriceFrom  string `query:"priceFrom"`
		BrandId    string `query:"brandId"`
		Sort       string `query:"sort"`
	}

	ProductResponse struct {
		Id          primitive.ObjectID     `bson:"_id"`
		Name        string                 `bson:"name"`
		Category    map[string]interface{} `bson:"category"`
		Brand       map[string]interface{} `bson:"brand"`
		Price       uint                   `bson:"price"`
		Description string                 `bson:"description"`
		Images      []string               `bson:"images"`
		Quantity    uint                   `bson:"quantity"`
		Rest        uint                   `bson:"rest"`
		Status      string                 `bson:"status"`
		CreatedAt   string                 `bson:"createdAt"`
		UpdatedAt   string                 `bson:"updatedAt"`
	}
)

func (p Product) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Required),
		validation.Field(&p.Description, validation.Required),
		validation.Field(&p.Images, validation.Length(1, 20)),
	)
}
