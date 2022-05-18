package route

import (
	"echo-app/config"
	"echo-app/controller"
	"echo-app/validation"
)

var (
	// controller
	userCtrl     = controller.User{}
	productCtrl  = controller.Product{}
	reviewCtrl   = controller.Review{}
	adminCtrl    = controller.Admin{}
	categoryCtrl = controller.Category{}
	brandCtrl    = controller.Brand{}

	//validate
	userVal     = validation.User{}
	productVal  = validation.Product{}
	reviewVal   = validation.Review{}
	adminVal    = validation.Admin{}
	categoryVal = validation.Category{}
	brandVal    = validation.Brand{}

	//secret key
	secretKeyBytes = []byte(config.GetEnv().Jwt.SecretKey)
)
