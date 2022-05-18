package controller

import (
	"echo-app/model"
	"echo-app/utils"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Review struct{}

func (r Review) GetListReview(c echo.Context) error {
	productId := c.Param("id")

	query := c.Get("query").(model.ReviewQuery)

	results, err := reviewService.GetListReview(productId, query)
	if err != nil {
		fmt.Println(err.Error())
		return utils.Response400(c, nil, utils.InvalidData)
	}
	return utils.Response200(c, results, utils.CreateSuccessFully)
}

func (r Review) CreateReview(c echo.Context) error {
	userId, err := utils.GetUserId(c)
	if err != nil {
		return utils.Response404(c, nil, utils.InvalidData)
	}

	payloadInterface := c.Get("body")
	payload, ok := payloadInterface.(model.CreateReview)
	if !ok {
		return utils.Response404(c, nil, utils.InvalidData)
	}

	productIdString := c.Param("id")
	productId, err := primitive.ObjectIDFromHex(productIdString)
	if err != nil {
		return utils.Response404(c, nil, utils.InvalidData)
	}

	result, err := reviewService.CreateReview(userId, productId, payload)
	if err != nil {
		return utils.Response400(c, "", err.Error())
	}
	return utils.Response200(c, result, utils.CreateSuccessFully)
}
