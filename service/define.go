package service

import "echo-app/dao"

var (
	userDAO      = dao.User{}
	orderDAO     = dao.Order{}
	orderItemDAO = dao.OrderItem{}
	cartDAO      = dao.Cart{}
	cartItem     = dao.CartItem{}
	productDAO = dao.Product{}
	reviewDAO  = dao.Review{}
)
