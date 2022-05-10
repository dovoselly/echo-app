package controllers

import (
	"echo-app/models"
	"echo-app/services"
	"echo-app/utils"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ListReview(c echo.Context) error {
	idString := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		return utils.Response404(c, nil, utils.InvalidData)
	}

	queryInterface := c.Get("query")
	query, ok := queryInterface.(models.ReviewQuery)
	if !ok {
		return utils.Response404(c, nil, utils.InvalidData)
	}

	results, err := services.ListReview(id, query)
	if err != nil {

	}
	return utils.Response200(c, results, utils.CreateSuccessFully)
}

func CreateReview(c echo.Context) error {
	payloadInterface := c.Get("payload")

	payload, ok := payloadInterface.(models.CreateReview)
	if !ok {
		return utils.Response404(c, nil, utils.InvalidData)
	}

	err := services.CreateReview(payload)
	if err != nil {
		return utils.Response200(c, "", err.Error())
	}
	return utils.Response200(c, "", utils.CreateSuccessFully)
}
