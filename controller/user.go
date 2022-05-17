package controller

import (
	"echo-app/models"
	"echo-app/services"
	"echo-app/utils"

	"github.com/labstack/echo/v4"
)

type User struct{}

func (u User) ChangePassword(c echo.Context) error {
	var (
		body = c.Get("body").(models.UserChangePassword)
	)

	id, _err := utils.GetUserId(c)
	if _err != nil {
		return _err
	}

	// process
	err := services.ChangeUserPassword(id, body)
	if err != nil {
		return utils.Response400(c, nil, err.Error())
	}

	return utils.Response200(c, id, "")
}

func (u User) GetInfo(c echo.Context) error {

	// Get Id in token
	ID, _err := utils.GetUserId(c)
	if _err != nil {
		return _err
	}

	// process
	info, err := services.GetUserInfo(ID)

	if err != nil {
		return utils.Response400(c, nil, err.Error())

	}

	return utils.Response200(c, info, "")
}

func (u User) UpdateInfo(c echo.Context) error {
	var (
		body = c.Get("body").(models.UserUpdate)
	)

	// Get Id in token
	ID, _err := utils.GetUserId(c)
	if _err != nil {
		return _err
	}

	err := services.UpdateUserInfo(ID, body)
	if err != nil {
		return utils.Response400(c, nil, err.Error())
	}
	return utils.Response200(c, ID, "")
}
