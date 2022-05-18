package route

import (
	"echo-app/controller"
	"echo-app/validation"
)

var (
	// controller
	userCtrl    = controller.User{}
	authCtrl    = controller.Auth{}
	orderCtrl   = controller.Order{}
	productCtrl = controller.Product{}

	//validate
	userVal    = validation.User{}
	productVal = validation.Product{}
)
