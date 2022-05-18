package controller

import (
	"echo-app/model"
	"echo-app/util"
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
		return util.Response400(c, nil, util.InvalidData)
	}
	return util.Response200(c, results, "")
}

func (r Review) CreateReview(c echo.Context) error {
	userId, err := util.GetUserId(c)
	if err != nil {
		return util.Response400(c, nil, util.InvalidData)
	}

	productId := c.Param("id")

	body := c.Get("body").(model.CreateReview)

	result, err := reviewService.CreateReview(userId, productId, body)
	if err != nil {
		return util.Response400(c, nil, util.InvalidData)
	}
	return util.Response200(c, result.InsertedID, util.CreateSuccessFully)
}
