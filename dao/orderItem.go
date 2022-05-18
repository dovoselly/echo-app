package dao

import (
	"context"
	"echo-app/database"
	"echo-app/model"
	"echo-app/utils"

	"go.mongodb.org/mongo-driver/bson"
)

type OrderItem struct{}

func (o OrderItem) GetByUserId() ([]model.OrderItemBSON, error) {
	var (
		orderItems   []model.OrderItemBSON
		orderItemCol = database.OrderItemCol()

		ctx = context.Background()
	)

	cursor, err := orderItemCol.Find(ctx, bson.M{})
	if err != nil {
		return orderItems, err
	}

	if err = cursor.All(context.Background(), &orderItems); err != nil {
		return orderItems, err
	}

	return orderItems, nil
}

func (o OrderItem) Create(body []model.OrderItemBSON) (string, error) {
	var data []interface{}
	for _, t := range body {
		data = append(data, t)
	}
	// create order
	_, err := database.OrderItemCol().InsertMany(utils.Ctx, data)
	if err != nil {
		return "", err
	}

	return "", nil
}
