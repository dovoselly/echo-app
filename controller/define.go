package controller

import "echo-app/service"

var (
	adminService    = service.Admin{}
	categoryService = service.Category{}
	brandService    = service.Brand{}
	userService    = service.User{}
	productService = service.Product{}
	reviewService  = service.Review{}
)
