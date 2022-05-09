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
		Items      []OrderItem        `bson:"items"`
		CreatedAt  time.Time          `bson:"createdAt"`
		UpdatedAt  time.Time          `bson:"updatedAt"`
	}

	OrderResponse struct {
		ID         string      `json:"_id"`
		UserId     User        `json:"userId"`
		DeliveryId string      `json:"deliveryId"`
		OrderCode  string      `json:"orderCode"`
		Status     string      `json:"status"`
		TotalPrice int64       `json:"totalPrice"`
		Note       string      `json:"note"`
		Payment    PaymentType `json:"payment"`
		Items      []OrderItem `json:"items"`
		CreatedAt  time.Time   `json:"createdAt"`
		UpdatedAt  time.Time   `json:"updatedAt"`
	}

	OrderCreate struct {
		UserId     string      `json:"userId"`
		DeliveryId string      `json:"deliveryId"`
		OrderCode  string      `json:"orderCode"`
		Status     string      `json:"status"`
		TotalPrice int64       `json:"totalPrice"`
		Note       string      `json:"note"`
		Payment    PaymentType `json:"payment"`
		Items      []OrderItem `json:"items"`
	}

	PaymentType struct {
		Method string `json:"method" bson:"method"`
		Status bool   `json:"status" bson:"status"`
	}
)
