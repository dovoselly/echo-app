package controllers

import (
	"echo-app/models"
	"echo-app/services"
	"echo-app/utils"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ListProduct(c echo.Context) error {
	//get query from middleware
	queryInterface := c.Get("query")
	query, ok := queryInterface.(models.ProductQuery)
	if !ok {
		return utils.Response404(c, nil, utils.InvalidData)
	}

	results, err := services.ListProduct(query)
	if err != nil {
		return utils.Response200(c, results, err.Error())
	}
	return utils.Response200(c, results, "")
}

func ProductDetail(c echo.Context) error {
	idString := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		return utils.Response404(c, nil, utils.InvalidData)
	}
	results, err := services.ProductDetail(id)
	if err != nil {
		return utils.Response200(c, results, err.Error())
	}
	return utils.Response200(c, results, "")
}
