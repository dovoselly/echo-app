package dao

import (
	"echo-app/database"
	"echo-app/model"
	"echo-app/utils"

	"go.mongodb.org/mongo-driver/bson"
)

type OrderItem struct{}

func (o OrderItem) GetByUserId() ([]model.OrderItemBSON, error) {
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
