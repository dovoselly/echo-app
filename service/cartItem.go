package service

import (
	"echo-app/model"
	"fmt"
)

type CartItem struct{}

func (c CartItem) Create(body []model.CartItemBSON) error {
	fmt.Println("CHECK TAI SERVICE CART_ITEM !!!!!")
	return cartItemDAO.Create(body)
}
