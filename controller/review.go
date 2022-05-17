package controller

import (
	"echo-app/model"
	"echo-app/service"
	"echo-app/util"
	"fmt"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ListReview(c echo.Context) error {
	productIdString := c.Param("id")
	productId, err := primitive.ObjectIDFromHex(productIdString)
	fmt.Println(productIdString)
	if err != nil {
		return util.Response400(c, nil, util.InvalidData)
	}

	queryInterface := c.Get("query")
	query, ok := queryInterface.(model.ReviewQuery)
	if !ok {
		return util.Response400(c, nil, util.InvalidData)
	}

	results, err := service.ListReview(productId, query)
	if err != nil {
		return util.Response400(c, nil, util.InvalidData)
	}
	return util.Response200(c, results, util.CreateSuccessFully)
}

func CreateReview(c echo.Context) error {
	userId, err := util.GetUserId(c)
	if err != nil {
		return util.Response404(c, nil, util.InvalidData)
	}

	payloadInterface := c.Get("body")
	payload, ok := payloadInterface.(model.CreateReview)
	if !ok {
		return util.Response404(c, nil, util.InvalidData)
	}

	productIdString := c.Param("id")
	productId, err := primitive.ObjectIDFromHex(productIdString)
	if err != nil {
		return util.Response404(c, nil, util.InvalidData)
	}

	err = service.CreateReview(userId, productId, payload)
	if err != nil {
		return util.Response200(c, "", err.Error())
	}
	return util.Response200(c, "", util.CreateSuccessFully)
}
