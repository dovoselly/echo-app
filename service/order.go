package service

import (
	"echo-app/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct{}

func (o Order) GetByUserId(id string) ([]model.OrderResponse, error) {
	var (
		orders = make([]model.OrderResponse, 0)
	)

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return orders, err
	}

	// get orders in db
	ordersBSON, err := orderDAO.GetByUserId(objId)
	if err != nil {
		return orders, err
	}

	// Convert to OrderJSON
	for _, orderBSON := range ordersBSON {
		orderItems := make([]model.OrderItemResponse, 0)

		// convert items in order
		for _, v := range orderBSON.Items {
			orderItem := v.ConvertToJSON()
			orderItems = append(orderItems, orderItem)
		}

		orderJSON := orderBSON.ConvertToJSON(orderItems)
		orders = append(orders, orderJSON)
	}

	return orders, nil
}

func (o Order) Create(idUser string, body model.OrderCreate) (string, error) {
	var (
		orderBSON model.OrderCreateBSON
	)

	objId, err := primitive.ObjectIDFromHex(idUser)
	if err != nil {
		return "", err
	}

	// insert to order-items db
	listItemJson := make([]model.OrderItemBSON, 0)
	for _, v := range body.Items {
		listItemJson = append(listItemJson, v.ConvertToBSON())
	}

	if _, err := orderItemDAO.Create(listItemJson); err != nil {
		return "", err
	}

	// get list id insert to order db
	ListIdItemJson := make([]primitive.ObjectID, 0)
	for _, v := range listItemJson {
		ListIdItemJson = append(ListIdItemJson, v.Id)
	}

	// convert orderCreate to orderCreateBson
	orderBSON = body.ConvertToBSON(ListIdItemJson, objId)

	// create
	orderId, err := orderDAO.CreateOrder(orderBSON)
	if err != nil {
		return "", err
	}

	return orderId, nil
}
