package controller

import (
	"echo-app/model"
	"echo-app/util"
	"fmt"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct{}

func (p Product) GetListProduct(c echo.Context) error {
	//get query from middleware
	query := c.Get("query").(model.ProductQuery)

	results, err := productService.GetListProduct(query)
	if err != nil {
		fmt.Println(err.Error())
		return util.Response400(c, nil, util.InvalidData)
	}
	return util.Response200(c, results, "")
}

func (p Product) GetProductDetail(c echo.Context) error {
	id := c.Param("id")

	results, err := productService.GetProductDetail(id)
	if err != nil {
		return util.Response400(c, nil, util.InvalidData)
	}
	return util.Response200(c, results, "")
}

func (p Product) Create(c echo.Context) (err error) {
	var (
		body = c.Get("body").(model.ProductCreate)
	)

	productID, err := productService.Create(body)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, bson.M{"_id": productID}, "")
}

func (p Product) Update(c echo.Context) (err error) {
	var (
		body = c.Get("body").(model.ProductUpdate)
		id   = c.Get("id").(string)
	)

	objID, _ := primitive.ObjectIDFromHex(id)

	if err := productService.Update(objID, body); err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, bson.M{"_id": id}, "")
}

func (p Product) UpdateStatus(c echo.Context) error {
	var (
		id = c.Get("id").(string)
	)

	objID, _ := primitive.ObjectIDFromHex(id)
	//process
	if err := productService.UpdateStatus(objID); err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, bson.M{"_id": id}, "")
}
