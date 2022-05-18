package controller

import (
	"echo-app/model"
	"echo-app/service"
	"echo-app/util"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func Register(c echo.Context) error {

	var (
		body = c.Get("body").(model.UserRegister)
	)

	// Process data
	rawData, err := service.UserRegister(body)

	if err != nil {

		return util.Response400(c, nil, err.Error())
	}

	// Success
	return util.Response200(c, bson.M{
		"_id":       rawData.ID,
		"createdAt": rawData.CreatedAt,
	}, "")
}

func Login(c echo.Context) error {
	var (
		user = c.Get("body").(model.UserLogin)
	)

	// process data
	token, err := service.Login(user)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	data := map[string]interface{}{
		"token": token,
	}

	return util.Response200(c, data, "")
}
