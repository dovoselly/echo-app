package service

import (
	"echo-app/model"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cart struct{}

func (c Cart) Create(userId string, body model.CartCreate) (string, error) {
	var (
		cartBson       model.CartCreateBSON
		cartItemsBson  = make([]model.CartItemBSON, 0)
		listIdCartItem = make([]primitive.ObjectID, 0)
	)
	// convert to objID
	objId, _ := primitive.ObjectIDFromHex(userId)

	// Convert CartCreate to CartBson
	// get cartItems bson
	for _, v := range body.Items {
		cartItemsBson = append(cartItemsBson, v.ConvertToBSON())
	}

	// creat many cart item
	if err := c.createCartItems(cartItemsBson); err != nil {
		return "", err
	}

	// get list id item
	for _, v := range cartItemsBson {
		listIdCartItem = append(listIdCartItem, v.Id)
	}

	// convert body cart
	cartBson = body.ConvertToBSON(listIdCartItem, objId)

	// create
	cartId, err := cartDAO.Create(cartBson)
	if err != nil {
		return "", nil
	}

	return cartId, nil
}

func (c Cart) createCartItems(cartItemsBson []model.CartItemBSON) error {
	fmt.Println("CHECK TAI SERVICE CART !!!!!")

	return cartItemDAO.Create(cartItemsBson)
}
