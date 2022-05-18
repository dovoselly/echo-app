package route

import (
	"echo-app/controller"
	"echo-app/validation"
)

var (
	// controller
	userCtrl     = controller.User{}
	adminCtrl    = controller.Admin{}
	categoryCtrl = controller.Category{}
	brandCtrl    = controller.Brand{}

	//validate
	userVal     = validation.User{}
	adminVal    = validation.Admin{}
	categoryVal = validation.Category{}
	brandVal    = validation.Brand{}
)
