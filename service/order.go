package service

import (
	"echo-app/dao"
	"echo-app/model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllOrderByUserId(ID primitive.ObjectID) ([]model.OrderResponse, error) {

	orders := make([]model.OrderResponse, 0)

	// get orders in db
	ordersBSON, err := dao.GetAllOrdersByUserId(ID)
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

func CreateOrder(ID primitive.ObjectID, body model.OrderCreate) error {

	// get list item insert to order-items db
	listItemJson := make([]model.OrderItemBSON, 0)
	for _, v := range body.Items {
		listItemJson = append(listItemJson, v.ConvertToOrderItemBSON())
	}

	if err := dao.CreateOrderItems(listItemJson); err != nil {
		return err
	}

	// get list id insert to order db
	ListIdItemJson := make([]primitive.ObjectID, 0)
	for _, v := range listItemJson {
		ListIdItemJson = append(ListIdItemJson, v.ID)
	}

	body.Status = "PENDING"
	orderBSON := model.OrderCreateBSON{
		ID:         primitive.NewObjectID(),
		UserId:     ID,
		DeliveryId: body.DeliveryId,
		OrderCode:  body.OrderCode,
		Status:     body.Status,
		TotalPrice: body.TotalPrice,
		Note:       body.Note,
		Payment:    body.Payment,
		Items:      ListIdItemJson,
		CreatedAt:  time.Now(),
	}

	err := dao.CreateOrder(orderBSON)
	if err != nil {
		return err
	}

	return nil
}
