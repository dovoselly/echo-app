package dao

import (
	"context"
	"echo-app/database"
	"echo-app/model"

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

func (o OrderItem) CreateOrderItems(body []model.OrderItemBSON) error {
	var (
		orderItemCol = database.OrderItemCol()
		ctx          = context.Background()
	)

	var data []interface{}
	for _, t := range body {
		data = append(data, t)
	}
	// create order
	_, err := orderItemCol.InsertMany(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
