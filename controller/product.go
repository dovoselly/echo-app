package controller

import (
	"echo-app/model"
	"echo-app/service"
	"echo-app/util"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProduct(c echo.Context) error {
	return c.JSON(http.StatusOK, "create products")
}
func ListProduct(c echo.Context) error {
	//get query from middleware
	queryInterface := c.Get("query")
	query, ok := queryInterface.(model.ProductQuery)
	if !ok {
		return util.Response404(c, nil, util.InvalidData)
	}

	results, err := service.ListProduct(query)
	if err != nil {
		return util.Response200(c, results, err.Error())
	}
	return util.Response200(c, results, "")
}

func ProductDetail(c echo.Context) error {
	idString := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		return util.Response404(c, nil, util.InvalidData)
	}
	results, err := service.ProductDetail(id)
	if err != nil {
		return util.Response200(c, results, err.Error())
	}
	return util.Response200(c, results, "")
}
