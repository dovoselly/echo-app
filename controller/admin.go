package controller

import (
	"echo-app/model"
	"echo-app/util"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Admin struct{}

func (a Admin) AdminLogin(c echo.Context) error {
	var admin = c.Get("body").(model.AdminLoginBody)

	// process data
	token, err := adminService.AdminLogin(admin)

	// if error
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// token
	data := map[string]interface{}{
		"token":   token,
		"isAdmin": true,
	}
	return util.Response200(c, data, "")
}

func (a Admin) MyProfileAdmin(c echo.Context) error {
	// jwtPayload get id
	jwtPayload, _ := util.GetJWTPayload(c)
	// admin id
	adminID := jwtPayload["id"].(string)

	// get admin profile
	profile, err := adminService.GetMyProfile(adminID)

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	data := map[string]interface{}{
		"profile": profile,
	}

	//success
	return util.Response200(c, data, "")
}

func (a Admin) UpdateMyProfileAdmin(c echo.Context) error {
	var body = c.Get("body").(model.Admin)

	// jwtPayload for get id
	jwtPayload, _ := util.GetJWTPayload(c)
	id := jwtPayload["id"].(string)

	// UpdateProfile
	err := adminService.UpdateMyProfile(id, body)

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, id, "")
}

func (a Admin) ChangePasswordAdmin(c echo.Context) error {
	return c.JSON(http.StatusOK, "Change password admin")
}

func (a Admin) ChangeAvatarAdmin(c echo.Context) error {
	return c.JSON(http.StatusOK, "Change avatar admin")
}
