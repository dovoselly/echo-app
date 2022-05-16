package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	OrderBSON struct {
		ID         primitive.ObjectID `bson:"_id"`
		UserId     primitive.ObjectID `bson:"userId"`
		DeliveryId primitive.ObjectID `bson:"deliveryId"`
		OrderCode  string             `bson:"orderCode"`
		Status     string             `bson:"status"`
		TotalPrice int64              `bson:"totalPrice"`
		Note       string             `bson:"note"`
		Payment    PaymentType        `bson:"payment"`
		Items      []OrderItemBSON    `bson:"items"`
		CreatedAt  time.Time          `bson:"createdAt"`
		UpdatedAt  time.Time          `bson:"updatedAt"`
	}
	OrderCreateBSON struct {
		ID         primitive.ObjectID   `bson:"_id"`
		UserId     primitive.ObjectID   `bson:"userId"`
		DeliveryId primitive.ObjectID   `bson:"deliveryId"`
		OrderCode  string               `bson:"orderCode"`
		Status     string               `bson:"status"`
		TotalPrice int64                `bson:"totalPrice"`
		Note       string               `bson:"note"`
		Payment    PaymentType          `bson:"payment"`
		Items      []primitive.ObjectID `bson:"items"`
		CreatedAt  time.Time            `bson:"createdAt"`
		UpdatedAt  time.Time            `bson:"updatedAt"`
	}

	OrderResponse struct {
		ID         primitive.ObjectID  `json:"_id"`
		UserId     primitive.ObjectID  `json:"userId"`
		DeliveryId primitive.ObjectID  `json:"deliveryId"`
		OrderCode  string              `json:"orderCode"`
		Status     string              `json:"status"`
		TotalPrice int64               `json:"totalPrice"`
		Note       string              `json:"note"`
		Payment    PaymentType         `json:"payment"`
		Items      []OrderItemResponse `json:"items"`
		CreatedAt  time.Time           `json:"createdAt"`
		UpdatedAt  time.Time           `json:"updatedAt"`
	}

	OrderCreate struct {
		UserId     primitive.ObjectID `json:"userId"`
		DeliveryId primitive.ObjectID `json:"deliveryId"`
		OrderCode  string             `json:"orderCode"`
		Status     string             `json:"status"`
		TotalPrice int64              `json:"totalPrice"`
		Note       string             `json:"note"`
		Payment    PaymentType        `json:"payment"`
		Items      []OrderItemCreate  `json:"items"`
	}

	PaymentType struct {
		Method string `json:"method" bson:"method"`
		Status bool   `json:"status" bson:"status"`
	}
)
