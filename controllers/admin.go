package controllers

import (
	"echo-app/models"
	"echo-app/services"
	"echo-app/utils"
	"github.com/labstack/echo/v4"
)

func AdminLogin(c echo.Context) error {
	var admin = c.Get("body").(models.AdminLoginBody)

	// process data
	token, err := services.AdminLogin(admin)

	// if error
	if err != nil {
		return utils.Response400(c, nil, err.Error())
	}

	// token
	data := map[string]interface{}{
		"token":   token,
		"isAdmin": true,
	}
	return utils.Response200(c, data, "")
}
