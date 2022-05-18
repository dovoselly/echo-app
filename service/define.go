package service

import "echo-app/dao"

var (
	adminDAO     = dao.Admin{}
	userDAO      = dao.User{}
	orderDAO     = dao.Order{}
	orderItemDAO = dao.OrderItem{}
	categoryDao  = dao.Category{}
	brandDao     = dao.Brand{}
	//cartDAO      = dao.Cart{}
	//cartItem     = dao.CartItem{}
	productDAO = dao.Product{}
	reviewDAO  = dao.Review{}
)
