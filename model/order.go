package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Order struct {
		ID            primitive.ObjectID `json:"_id" bson:"_id"`
		UserId        primitive.ObjectID `json:"userId" bson:"userId"`
		DeliveryId    primitive.ObjectID `json:"deliveryId" bson:"delivery"`
		OrderCode     string             `json:"orderCode" bson:"orderCode"`
		Status        string             `json:"status" bson:"status"`
		PaymentMethod string             `json:"paymentMethod" bson:"paymentMethod"`
		PaymentStatus string             `json:"paymentStatus" bson:"paymentStatus"`
		TotalPrice    int64              `json:"totalPrice" bson:"totalPrice"`
		Note          string             `json:"note" bson:"note"`
		Items         []OrderItem        `json:"items" bson:"items"`
		CreatedAt     string             `json:"createdAt" bson:"createdAt"`
		UpdatedAt     string             `json:"updatedAt" bson:"updatedAt"`
	}

	OrderItem struct {
		ID        primitive.ObjectID `json:"_id" bson:"_id"`
		ProductId primitive.ObjectID `json:"productId" bson:"productId"`
		Quantity  int16              `json:"quantity" bson:"quantity"`
		Note      string             `json:"note" bson:"note"`
		Price     int64              `json:"price" bson:"price"`
		CreatedAt string             `json:"createdAt" bson:"createdAt"`
		UpdatedAt string             `json:"updatedAt" bson:"updatedAt"`
	}
)
