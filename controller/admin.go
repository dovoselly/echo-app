package controller

import (
	"echo-app/model"
	"echo-app/service"
	"echo-app/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AdminLogin(c echo.Context) error {
	var admin = c.Get("adminLoginBody").(model.AdminLoginBody)

	// process data
	token, err := service.AdminLogin(admin)

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

func MyProfileAdmin(c echo.Context) error {
	// jwtPayload get id
	jwtPayload, _ := utils.GetJWTPayload(c)
	// admin id
	adminID := jwtPayload["id"].(string)

	// get admin profile
	profile, err := service.MyProfileAdmin(adminID)

	// if err
	if err != nil {
		return utils.Response400(c, nil, err.Error())
	}

	data := map[string]interface{}{
		"profile": profile,
	}

	//success
	return utils.Response200(c, data, "")
}

func UpdateMyProfileAdmin(c echo.Context) error {
	var body = c.Get("adminRequestBody").(model.Admin)

	// jwtPayload for get id
	jwtPayload, _ := utils.GetJWTPayload(c)
	id := jwtPayload["id"].(string)

	// UpdateProfile
	err := service.UpdateMyProfileAdmin(id, body)

	// if err
	if err != nil {
		return utils.Response400(c, nil, err.Error())
	}

	return utils.Response200(c, id, "")
}

func ChangePasswordAdmin(c echo.Context) error {
	return c.JSON(http.StatusOK, "Change password admin")
}

func ChangeAvatarAdmin(c echo.Context) error {
	return c.JSON(http.StatusOK, "Change avatar admin")
}
