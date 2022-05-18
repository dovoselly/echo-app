package controller

import (
	"echo-app/model"
	"echo-app/utils"
	"fmt"
	"github.com/labstack/echo/v4"
)

type Review struct{}

func (r Review) GetListReview(c echo.Context) error {
	ID := c.Param("id")

	query := c.Get("query").(model.ReviewQuery)

	results, err := reviewService.GetListReview(ID, query)
	if err != nil {
		fmt.Println(err.Error())
		return utils.Response400(c, nil, utils.InvalidData)
	}
	return utils.Response200(c, results, "")
}

func (r Review) CreateReview(c echo.Context) error {
	userId, err := utils.GetUserId(c)
	if err != nil {
		return utils.Response400(c, nil, utils.InvalidData)
	}

	productId := c.Param("id")

	body := c.Get("body").(model.CreateReview)

	result, err := reviewService.CreateReview(userId, productId, body)
	if err != nil {
		return utils.Response400(c, nil, utils.InvalidData)
	}
	return utils.Response200(c, result.InsertedID, utils.CreateSuccessFully)
}
