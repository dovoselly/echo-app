package dao

import (
	"context"
	"echo-app/database"
	"echo-app/model"
	"echo-app/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct{}

func (o Order) GetAllByUserId(ID primitive.ObjectID) ([]model.OrderBSON, error) {

	var (
		orders   []model.OrderBSON
		orderCol = database.OrderCol()
		ctx      = context.Background()
	)

	pipeline := []bson.M{
		{
			"$match": bson.M{"userId": ID},
		},
		{"$lookup": bson.M{
			"from":         "orderItems",
			"localField":   "items",
			"foreignField": "_id",
			"as":           "items",
		}},
	}

	cursor, err := orderCol.Aggregate(ctx, pipeline)

	err = cursor.All(context.Background(), &orders)

	if err != nil {
		return orders, err
	}

	return orders, nil
}

func (o Order) CreateOrder(body model.OrderCreateBSON) (primitive.ObjectID, error) {
	// create order
	result, err := database.OrderCol().InsertOne(utils.Ctx, body)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}
