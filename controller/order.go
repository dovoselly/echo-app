package controller

import (
	"echo-app/model"
	"echo-app/utils"

	"github.com/labstack/echo/v4"
)

type Order struct{}

func (o Order) GetByUserId(c echo.Context) error {
	// Get Id in token
	ID, err := utils.GetUserId(c)
	if err != nil {
		return err
	}

	data, err := orderService.GetByUserId(ID)
	if err != nil {
		return utils.Response400(c, nil, err.Error())
	}

	return utils.Response200(c, data, "")
}

func (o Order) CreateOrder(c echo.Context) error {
	var (
		body = c.Get("body").(model.OrderCreate)
	)

	// get id user in token
	idUser, err := utils.GetUserId(c)
	if err != nil {
		return err
	}

	// process
	id, err := orderService.CreateOrder(idUser, body)
	if err != nil {
		return utils.Response400(c, nil, err.Error())
	}

	return utils.Response200(c, id, "")
}
