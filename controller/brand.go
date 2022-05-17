package controller

import (
	"echo-app/models"
	"echo-app/service"
	"echo-app/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateBrand(c echo.Context) error {
	var body = c.Get("body").(models.BrandCreateBody)

	// process data
	err := service.CreateBrand(body)

	// if err
	if err != nil {
		return utils.Response400(c, nil, err.Error())
	}

	return utils.Response200(c, body, "")
}

func GetListBrand(c echo.Context) error {
	brands, err := service.GetListBrand()
	if err != nil {
		return utils.Response400(c, nil, err.Error())
	}
	return utils.Response200(c, brands, "")
}

func GetBrandByID(c echo.Context) error {
	var strID = c.Get("id").(string)

	// process
	brand, err := service.GetBrandByID(strID)

	// if error
	if err != nil {
		return utils.Response400(c, nil, err.Error())
	}

	return utils.Response200(c, brand, "")
}

func UpdateBrand(c echo.Context) error {
	var (
		ID   = c.Get("id").(string)
		body = c.Get("body").(models.BrandUpdateBody)
	)

	// process data
	err := service.UpdateBrandByID(ID, body)
	if err != nil {
		return utils.Response400(c, nil, err.Error())
	}

	return utils.Response200(c, nil, "")

}

func DeleteBrandByID(c echo.Context) error {
	var id = c.Get("id").(string)

	//process
	err := service.DeleteBrandByID(id)
	if err != nil {
		return utils.Response400(c, nil, err.Error())
	}

	return utils.Response200(c, nil, "")
}

func DisabledBrand(c echo.Context) error {
	return c.JSON(http.StatusOK, "Disabled brand")
}
func EnabledBrand(c echo.Context) error {
	return c.JSON(http.StatusOK, "Enabled brand")
}
