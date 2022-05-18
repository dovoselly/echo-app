package controller

import (
	"echo-app/model"
	"echo-app/util"

	"github.com/labstack/echo/v4"
)

type Order struct{}

func (o Order) GetByUserId(c echo.Context) error {
	// Get Id in token
	id, err := util.GetUserId(c)
	if err != nil {
		return err
	}

	data, err := orderService.GetByUserId(id)
	if err != nil {
		return util.Response400(c, nil, util.InvalidData)
	}

	return util.Response200(c, data, "")
}

func (o Order) Create(c echo.Context) error {
	var (
		body = c.Get("body").(model.OrderCreate)
	)

	// get id user in token
	objId, err := util.GetUserId(c)
	if err != nil {
		return err
	}

	// process
	id, err := orderService.Create(objId, body)
	if err != nil {
		return util.Response400(c, nil, util.InvalidData)
	}

	return util.Response200(c, id, "")
}
