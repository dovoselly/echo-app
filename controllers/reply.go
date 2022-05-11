package controllers

import (
	"echo-app/models"
	"echo-app/services"
	"echo-app/utils"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateReply(c echo.Context) error {
	idString := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		return utils.Response404(c, nil, utils.InvalidData)
	}

	payloadInterface := c.Get("payload")
	payload := payloadInterface.(models.CreateReply)

	err = services.CreateReply(id, payload)
	return utils.Response200(c, nil, err.Error())
}
