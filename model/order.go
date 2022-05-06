package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Order struct {
		ID         primitive.ObjectID `json:"_id" bson:"_id"`
		UserId     primitive.ObjectID `json:"userId" bson:"userId"`
		DeliveryId primitive.ObjectID `json:"deliveryId" bson:"delivery"`
		OrderCode  string             `json:"orderCode" bson:"orderCode"`
		Status     string             `json:"status" bson:"status"`
		Payment    PaymentType        `json:"payment" bson:"payment"`
		TotalPrice int64              `json:"totalPrice" bson:"totalPrice"`
		Note       string             `json:"note" bson:"note"`
		Items      []OrderItem        `json:"items" bson:"items"`
		CreatedAt  string             `json:"createdAt" bson:"createdAt"`
		UpdatedAt  string             `json:"updatedAt" bson:"updatedAt"`
	}

	PaymentType struct {
		Method string `json:"method" bson:"method"`
		Status bool   `json:"status" bson:"status"`
	}
)
