package service

import "echo-app/dao"

var (
	userDAO      = dao.User{}
	orderDAO     = dao.Order{}
	orderItemDAO = dao.OrderItem{}
	productDAO   = dao.Product{}
	cartDAO      = dao.Cart{}
	cartItem     = dao.CartItem{}
	reviewDAO    = dao.Review{}
)
