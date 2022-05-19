package controller

import (
	"echo-app/model"
	"echo-app/util"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct{}

func (u User) ChangePassword(c echo.Context) error {
	var (
		body = c.Get("body").(model.UserChangePassword)
	)

	// get id user in token
	id, err := util.GetUserId(c)
	if err != nil {
		return err
	}

	// process
	if err := userService.ChangePassword(id, body); err != nil {
		return util.Response400(c, nil, util.InvalidData)
	}

	return util.Response200(c, bson.M{"id": id}, "")
}

func (u User) GetInfo(c echo.Context) error {
	// Get id user in token
	id, err := util.GetUserId(c)
	if err != nil {
		return err
	}

	// process
	info, err := userService.GetInfo(id)
	if err != nil {
		return util.Response400(c, nil, util.InvalidData)

	}

	return util.Response200(c, info, "")
}

func (u User) UpdateInfo(c echo.Context) error {
	var (
		body = c.Get("body").(model.UserUpdate)
	)

	// get ID user in token
	id, err := util.GetUserId(c)
	if err != nil {
		return err
	}

	//process
	if err := userService.UpdateInfo(id, body); err != nil {
		return util.Response400(c, nil, util.InvalidData)
	}

	return util.Response200(c, bson.M{"_id": id}, "")
}
