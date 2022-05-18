package controller

import "echo-app/service"

var (
	userService    = service.User{}
	productService = service.Product{}
	reviewService  = service.Review{}
)
