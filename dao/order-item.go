package dao

import (
	"context"
	"echo-app/database"
	"echo-app/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAllOrderItemsByUserId() ([]models.OrderItemBSON, error) {
	var (
		orderItems   []models.OrderItemBSON
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

func CreateOrderItems(body []models.OrderItemBSON) error {
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
