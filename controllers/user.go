package controllers

import (
	"echo-app/models"
	"echo-app/services"
	"echo-app/utils"

	"github.com/labstack/echo/v4"
)

func UserChangePassword(c echo.Context) error {
	var (
		body = c.Get("body").(models.UserChangePassword)
		id   = c.Get("id").(string)
	)

	// process
	err := services.UserChangePassword(id, body)
	if err != nil {
		return utils.Response400(c, nil, err.Error())
	}

	return utils.Response200(c, id, "")
}
