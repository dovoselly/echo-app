package service

import "echo-app/dao"

var (
	orderDAO     = dao.Order{}
	orderItemDAO = dao.OrderItem{}
	productDAO   = dao.Product{}
)
