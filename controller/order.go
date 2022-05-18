package controller

import (
	"echo-app/models"
	"echo-app/service"
	"echo-app/util"

	"github.com/labstack/echo/v4"
)

func CreateOrder(c echo.Context) error {

	var (
		body = c.Get("body").(models.OrderCreate)
	)

	// Get Id in token
	ID, _err := util.GetUserId(c)
	if _err != nil {
		return _err
	}

	err := service.CreateOrder(ID, body)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, nil, "")
}

func GetAllOrdersByUserId(c echo.Context) error {
	// Get Id in token
	ID, _err := util.GetUserId(c)
	if _err != nil {
		return _err
	}

	data, err := service.GetAllOrderByUserId(ID)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, data, "")
}
