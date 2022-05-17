package controller

import (
	"echo-app/model"
	"echo-app/service"
	"echo-app/utils"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct{}

func (p Product) ListProduct(c echo.Context) error {
	//get query from middleware
	var (
		s = service.Product{}
	)
	queryInterface := c.Get("query")
	query, ok := queryInterface.(model.ProductQuery)
	if !ok {
		return utils.Response404(c, nil, utils.InvalidData)
	}

	results, err := s.ListProduct(query)
	if err != nil {
		return utils.Response200(c, results, err.Error())
	}
	return utils.Response200(c, results, "")
}

func (p Product) ProductDetail(c echo.Context) error {
	var (
		s = service.Product{}
	)
	idString := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		return utils.Response404(c, nil, utils.InvalidData)
	}
	results, err := s.ProductDetail(id)
	if err != nil {
		return utils.Response200(c, results, err.Error())
	}
	return utils.Response200(c, results, "")
}
