package controllers

import (
	"echo-app/models"
	"echo-app/services"
	"echo-app/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

func AdminLogin(c echo.Context) error {
	var admin = c.Get("adminLoginBody").(models.AdminLoginBody)

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

//func MyProfileAdmin(c echo.Context) error {
//	// jwtPayload get id
//	jwtPayload, _ := middlewares.GetJWTPayload(c)
//
//	// admin id
//	adminID := jwtPayload["id"].(string)
//
//	// get admin profile
//	profile, err := services.MyProfileAdmin(adminID)
//
//	// if err
//	if err != nil {
//		return utils.Response400(c, nil, err.Error())
//	}
//
//	// get admin
//	admin, err := services.
//}

//func AdminLogin(c echo.Context) error {
//	return c.JSON(http.StatusOK, "Admin login")
//}

func MyProfileAdmin(c echo.Context) error {
	return c.JSON(http.StatusOK, "Get Admin profile")
}

func UpdateMyProfileAdmin(c echo.Context) error {
	return c.JSON(http.StatusOK, "Update admin profile")
}

func ChangePasswordAdmin(c echo.Context) error {
	return c.JSON(http.StatusOK, "Change password admin")
}

func ChangeAvatarAdmin(c echo.Context) error {
	return c.JSON(http.StatusOK, "Change avatar admin")
}
