package controllers

import (
	"echo-app/models"
	"echo-app/services"
	"echo-app/utils"

	"github.com/labstack/echo/v4"
)

func ChangeUserPassword(c echo.Context) error {
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

func GetUserInfo(c echo.Context) error {

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

func UpdateUserInfo(c echo.Context) error {
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
