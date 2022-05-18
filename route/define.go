package route

import (
	"echo-app/config"
	"echo-app/controller"
	"echo-app/validation"
)

var (
	// controller
	userCtrl    = controller.User{}
	authCtrl    = controller.Auth{}
	orderCtrl   = controller.Order{}
	productCtrl = controller.Product{}
	reviewCtrl  = controller.Review{}

	//validate
	userVal    = validation.User{}
	productVal = validation.Product{}
	reviewVal  = validation.Review{}

	//secret key
	secretKeyBytes = []byte(config.GetEnv().Jwt.SecretKey)
)
