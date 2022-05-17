package controller

import (
	"echo-app/model"
	"echo-app/service"
	"echo-app/util"

	"github.com/labstack/echo/v4"
)

type User struct{}

func (u User) ChangePassword(c echo.Context) error {
	var (
		body = c.Get("body").(model.UserChangePassword)
	)

	id, _err := util.GetUserId(c)
	if _err != nil {
		return _err
	}

	// process
	err := service.ChangeUserPassword(id, body)
	if err != nil {
		return util.Response400(c, nil, err.Error())
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
		return util.Response400(c, nil, err.Error())

	}

	return util.Response200(c, info, "")
}

func (u User) UpdateInfo(c echo.Context) error {
	var (
		body = c.Get("body").(model.UserUpdate)
	)

	// Get Id in token
	ID, _err := util.GetUserId(c)
	if _err != nil {
		return _err
	}

	err := service.UpdateUserInfo(ID, body)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}
	return util.Response200(c, ID, "")
}
