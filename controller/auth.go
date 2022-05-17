package controller

import (
	"echo-app/models"
	"echo-app/services"
	"echo-app/utils"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func Register(c echo.Context) error {

	var (
		body = c.Get("body").(models.UserRegister)
	)

	// Process data
	rawData, err := services.UserRegister(body)

	if err != nil {

		return utils.Response400(c, nil, err.Error())
	}

	// Success
	return utils.Response200(c, bson.M{
		"_id":       rawData.ID,
		"createdAt": rawData.CreatedAt,
	}, "")
}

func Login(c echo.Context) error {
	var (
		user = c.Get("body").(models.UserLogin)
	)

	// process data
	token, err := services.Login(user)
	if err != nil {
		return utils.Response400(c, nil, err.Error())
	}

	data := map[string]interface{}{
		"token": token,
	}

	return utils.Response200(c, data, "")
}
