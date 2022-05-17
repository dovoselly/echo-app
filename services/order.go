package services

import (
	"echo-app/dao"
	"echo-app/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllOrderByUserId(ID primitive.ObjectID) ([]models.OrderResponse, error) {

	orders := make([]models.OrderResponse, 0)

	// get orders in db
	ordersBSON, err := dao.GetAllOrdersByUserId(ID)
	if err != nil {
		return orders, err
	}

	// Convert to OrderJSON
	for _, orderBSON := range ordersBSON {

		orderItems := make([]models.OrderItemResponse, 0)

		for _, v := range orderBSON.Items {
			orderItem := models.OrderItemResponse{
				ID:        v.ID,
				ProductId: v.ProductId,
				Price:     v.Price,
				Quantity:  v.Quantity,
				Note:      v.Note,
			}
			orderItems = append(orderItems, orderItem)
		}

		orderJSON := models.OrderResponse{
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

func CreateOrder(ID primitive.ObjectID, body models.OrderCreate) error {

	// get list item insert to order-items db
	listItemJson := make([]models.OrderItemBSON, 0)
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
	orderBSON := models.OrderCreateBSON{
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
