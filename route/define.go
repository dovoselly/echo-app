package route

import (
	"echo-app/controller"
	"echo-app/validation"
)

var (
	// controller
	userCtrl  = controller.User{}
	adminCtrl = controller.Admin{}

	//validate
	userVal  = validation.User{}
	adminVal = validation.Admin{}
)
