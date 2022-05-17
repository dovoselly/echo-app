package route

import (
	"echo-app/controller"
	"echo-app/validation"
)

var (
	// controller
	userCtrl    = controller.User{}
	productCtrl = controller.Product{}

	//validate
	userVal    = validation.User{}
	productVal = validation.Product{}
)
