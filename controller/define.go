package controller

import "echo-app/service"

var (
	authService  = service.Auth{}
	userService  = service.User{}
	orderService = service.Order{}

	productService = service.Product{}
	reviewService  = service.Review{}
)
