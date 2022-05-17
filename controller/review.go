package controller

import (
	"echo-app/models"
	"echo-app/service"
	"echo-app/utils"
	"fmt"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ListReview(c echo.Context) error {
	productIdString := c.Param("id")
	productId, err := primitive.ObjectIDFromHex(productIdString)
	fmt.Println(productIdString)
	if err != nil {
		return utils.Response400(c, nil, utils.InvalidData)
	}

	queryInterface := c.Get("query")
	query, ok := queryInterface.(models.ReviewQuery)
	if !ok {
		return utils.Response400(c, nil, utils.InvalidData)
	}

	results, err := service.ListReview(productId, query)
	if err != nil {
		return utils.Response400(c, nil, utils.InvalidData)
	}
	return utils.Response200(c, results, utils.CreateSuccessFully)
}

func CreateReview(c echo.Context) error {
	userId, err := utils.GetUserId(c)
	if err != nil {
		return utils.Response404(c, nil, utils.InvalidData)
	}

	payloadInterface := c.Get("body")
	payload, ok := payloadInterface.(models.CreateReview)
	if !ok {
		return utils.Response404(c, nil, utils.InvalidData)
	}

	productIdString := c.Param("id")
	productId, err := primitive.ObjectIDFromHex(productIdString)
	if err != nil {
		return utils.Response404(c, nil, utils.InvalidData)
	}

	err = service.CreateReview(userId, productId, payload)
	if err != nil {
		return utils.Response200(c, "", err.Error())
	}
	return utils.Response200(c, "", utils.CreateSuccessFully)
}
