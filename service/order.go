package service

import (
	"echo-app/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct{}

func (o Order) GetByUserId(ID primitive.ObjectID) ([]model.OrderResponse, error) {

	orders := make([]model.OrderResponse, 0)

	// get orders in db
	ordersBSON, err := orderDAO.GetAllByUserId(ID)
	if err != nil {
		return orders, err
	}

	// Convert to OrderJSON
	for _, orderBSON := range ordersBSON {

		orderItems := make([]model.OrderItemResponse, 0)

		for _, v := range orderBSON.Items {
			orderItem := model.OrderItemResponse{
				ID:        v.ID,
				ProductId: v.ProductId,
				Price:     v.Price,
				Quantity:  v.Quantity,
				Note:      v.Note,
			}
			orderItems = append(orderItems, orderItem)
		}

		orderJSON := model.OrderResponse{
			ID:         orderBSON.ID,
			UserId:     orderBSON.UserId,
			DeliveryId: orderBSON.DeliveryId,
			OrderCode:  orderBSON.OrderCode,
			Status:     orderBSON.Status,
			TotalPrice: orderBSON.TotalPrice,
			Note:       orderBSON.Note,
			Payment:    orderBSON.Payment,
			Items:      orderItems,
		}

		orders = append(orders, orderJSON)
	}

	return orders, nil
}

func (o Order) CreateOrder(id primitive.ObjectID, body model.OrderCreate) (primitive.ObjectID, error) {
	var (
		orderBSON model.OrderCreateBSON
	)

	// insert to order-items db
	listItemJson := make([]model.OrderItemBSON, 0)
	for _, v := range body.Items {
		listItemJson = append(listItemJson, v.ConvertToOrderItemBSON())
	}

	if err := orderItemDAO.CreateOrderItems(listItemJson); err != nil {
		return id, err
	}

	// get list id insert to order db
	ListIdItemJson := make([]primitive.ObjectID, 0)
	for _, v := range listItemJson {
		ListIdItemJson = append(ListIdItemJson, v.ID)
	}

	// convert orderCreate to orderCreateBson
	orderBSON = body.ConvertToBSON(ListIdItemJson)

	// create
	id, err := orderDAO.CreateOrder(orderBSON)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return id, nil
}
