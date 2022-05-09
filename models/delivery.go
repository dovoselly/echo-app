package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	Delivery struct {
		ID               primitive.ObjectID `json:"_id" bson:"_id"`
		ShippingUnitName string             `json:"shippingUnitName" bson:"shippingUnitName"`
		ShipFee          int                `json:"shipFee" bson:"shipFee"`
		Address          string             `json:"address" bson:"address"`
		CreatedAt        string             `json:"createdAt" bson:"createdAt"`
		UpdatedAt        string             `json:"updatedAt" bson:"updatedAt"`
	}
)
