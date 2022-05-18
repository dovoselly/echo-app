package controller

import (
	"echo-app/models"
	"echo-app/service"
	"echo-app/util"

	"github.com/labstack/echo/v4"
)

type User struct{}

func (u User) ChangePassword(c echo.Context) error {
	var (
		body = c.Get("body").(models.UserChangePassword)
	)

	id, _err := util.GetUserId(c)
	if _err != nil {
		return _err
	}

	// process
	if err := userService.ChangePassword(id, body); err != nil {
		return util.Response400(c, nil, util.InvalidData)
	}

	return util.Response200(c, id, "")
}

func (u User) GetInfo(c echo.Context) error {

	// Get Id in token
	ID, _err := util.GetUserId(c)
	if _err != nil {
		return _err
	}

	// process
	info, err := service.GetUserInfo(ID)
	if err != nil {
		return util.Response400(c, nil, util.InvalidData)

	}

	return util.Response200(c, info, "")
}

func (u User) UpdateInfo(c echo.Context) error {
	var (
		body = c.Get("body").(models.UserUpdate)
	)

	// Get Id in token
	ID, _err := util.GetUserId(c)
	if _err != nil {
		return _err
	}

	//process
	if err := userService.UpdateInfo(id, body); err != nil {
		return util.Response400(c, nil, util.InvalidData)
	}
	return util.Response200(c, ID, "")
}
