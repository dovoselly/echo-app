package controller

import (
	"echo-app/model"
	"echo-app/util"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct{}

func (o Order) GetByUserId(c echo.Context) error {
	// Get Id in token
	id, err := util.GetUserId(c)
	if err != nil {
		return err
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	data, err := orderService.GetByUserId(objID)
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
	UserID, err := util.GetUserId(c)
	if err != nil {
		return err
	}

	objID, err := primitive.ObjectIDFromHex(UserID)
	if err != nil {
		return err
	}

	// process
	orderID, err := orderService.Create(objID, body)
	if err != nil {
		return util.Response400(c, nil, util.InvalidData)
	}

	return util.Response200(c, bson.M{"_id": orderID}, "")
}
