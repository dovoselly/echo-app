package controller

import (
	"echo-app/model"
	"echo-app/util"
	"fmt"

	"github.com/labstack/echo/v4"
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
