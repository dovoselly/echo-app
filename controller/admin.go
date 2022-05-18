package controller

import (
	"echo-app/model"
	"echo-app/util"

	"github.com/labstack/echo/v4"
)

type Admin struct{}

func (Admin) AdminLogin(c echo.Context) error {
	var admin = c.Get("adminLoginBody").(model.AdminLoginBody)

	// process data
	token, err := adminService.AdminLogin(admin)

	// if error
	if err != nil {
		return util.Response400(c, nil, util.InvalidData)
	}

	// token
	data := map[string]interface{}{
		"token":   token,
		"isAdmin": true,
	}
	return util.Response200(c, data, "")
}

func (Admin) GetMyProfile(c echo.Context) error {
	// jwtPayload get id
	jwtPayload, _ := util.GetJWTPayload(c)
	// admin id
	adminID := jwtPayload["id"].(string)

	// get admin profile
	profile, err := adminService.GetMyProfile(adminID)

	// if err
	if err != nil {
		return util.Response400(c, nil, util.InvalidData)
	}

	data := map[string]interface{}{
		"profile": profile,
	}

	//success
	return util.Response200(c, data, "")
}

func (Admin) UpdateMyProfile(c echo.Context) error {
	var body = c.Get("adminRequestBody").(model.Admin)

	// jwtPayload for get id
	jwtPayload, _ := util.GetJWTPayload(c)
	id := jwtPayload["id"].(string)

	// UpdateProfile
	err := adminService.UpdateMyProfile(id, body)

	// if err
	if err != nil {
		return util.Response400(c, nil, util.InvalidData)
	}

	return util.Response200(c, id, "")
}

//func AdminLogin(c echo.Context) error {
//	return c.JSON(http.StatusOK, "Admin login")
//}

//func MyProfileAdmin(c echo.Context) error {
//	return c.JSON(http.StatusOK, "Get Admin profile")
//}
