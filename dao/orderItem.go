package dao

import (
	"echo-app/database"
	"echo-app/model"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAllOrderItemsByUserId() ([]model.OrderItemBSON, error) {
	var (
		orderItems []model.OrderItemBSON
	)

	cursor, err := database.OrderItemCol().Find(utils.Ctx, bson.M{})
	if err != nil {
		return orderItems, err
	}

	if err = cursor.All(utils.Ctx, &orderItems); err != nil {
		return orderItems, err
	}

	return orderItems, nil
}

func CreateOrderItems(body []model.OrderItemBSON) error {
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
