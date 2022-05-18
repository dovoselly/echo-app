package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	ProductBSON struct {
		Id          primitive.ObjectID `bson:"_id"`
		Name        string             `bson:"name"`
		CategoryId  primitive.ObjectID `bson:"categoryId"`
		BrandId     primitive.ObjectID `bson:"brandId"`
		Price       uint               `bson:"price"`
		Description string             `bson:"description"`
		Images      []string           `bson:"images"`
		Quantity    uint               `bson:"quantity"`
		Rest        uint               `bson:"rest"`
		Status      string             `bson:"status"`
		CreatedAt   time.Time          `bson:"createdAt"`
		UpdatedAt   time.Time          `bson:"updatedAt"`
	}

	ProductQuery struct {
		Page       int64  `query:"page"`
		Name       string `query:"name"`
		CategoryId string `query:"categoryId"`
		PriceFrom  string `query:"priceFrom"`
		BrandId    string `query:"brandId"`
		Sort       string `query:"sort"`
	}

	ProductResponse struct {
		Id          primitive.ObjectID     `json:"_id"`
		Name        string                 `json:"name"`
		Category    map[string]interface{} `json:"category"`
		Brand       map[string]interface{} `json:"brand"`
		Price       uint                   `json:"price"`
		Description string                 `json:"description"`
		Images      []string               `json:"images"`
		Quantity    uint                   `json:"quantity"`
		Rest        uint                   `json:"rest"`
		Status      string                 `json:"status"`
		CreatedAt   string                 `json:"createdAt"`
		UpdatedAt   string                 `json:"updatedAt"`
	}
)

func (p ProductBSON) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Required),
		validation.Field(&p.Description, validation.Required),
		validation.Field(&p.Images, validation.Length(1, 20)),
	)
}
