package controller

import (
	"echo-app/model"
	"echo-app/util"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Brand struct{}

func (Brand) Create(c echo.Context) error {
	var body = c.Get("body").(model.BrandCreateBody)

	// process data
	err := brandService.Create(body)

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, body, "")
}

func (Brand) GetList(c echo.Context) error {
	brands, err := brandService.GetList()
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}
	return util.Response200(c, brands, "")
}

func (Brand) GetByID(c echo.Context) error {
	var strID = c.Get("id").(string)

	// process
	brand, err := brandService.GetByID(strID)

	// if error
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, brand, "")
}

func (Brand) UpdateByID(c echo.Context) error {
	var (
		ID   = c.Get("id").(string)
		body = c.Get("body").(model.BrandUpdateBody)
	)

	// process data
	err := brandService.UpdateByID(ID, body)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, nil, "")

}

func (Brand) DeleteByID(c echo.Context) error {
	var id = c.Get("id").(string)

	//process
	err := brandService.DeleteByID(id)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, nil, "")
}

func (Brand) Disabled(c echo.Context) error {
	return c.JSON(http.StatusOK, "Disabled brand")
}
func (Brand) Enabled(c echo.Context) error {
	return c.JSON(http.StatusOK, "Enabled brand")
}
