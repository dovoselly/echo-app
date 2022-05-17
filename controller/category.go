package controller

import (
	"echo-app/models"
	"echo-app/service"
	"echo-app/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateCategory(c echo.Context) error {
	var body = c.Get("body").(models.CategoryCreateBody)

	// process data
	err := service.CreateCategory(body)

	// if err
	if err != nil {
		return utils.Response400(c, nil, err.Error())
	}

	return utils.Response200(c, body, "")
}

func GetListCategory(c echo.Context) error {
	categories, err := service.GetListCategory()
	if err != nil {
		return utils.Response400(c, nil, err.Error())
	}
	return utils.Response200(c, categories, "")
}

func GetCategoryByID(c echo.Context) error {
	var strID = c.Get("id").(string)

	// process
	category, err := service.GetCategoryByID(strID)

	// if error
	if err != nil {
		return utils.Response400(c, nil, err.Error())
	}

	return utils.Response200(c, category, "")
}

func UpdateCategoryByID(c echo.Context) error {
	var (
		ID   = c.Get("id").(string)
		body = c.Get("body").(models.CategoryUpdateBody)
	)

	// process data
	err := service.UpdateCategoryByID(ID, body)
	if err != nil {
		return utils.Response400(c, nil, err.Error())
	}

	return utils.Response200(c, nil, "")

}

func DeleteCategoryByID(c echo.Context) error {
	var id = c.Get("id").(string)

	//process
	err := service.DeleteCategoryByID(id)
	if err != nil {
		return utils.Response400(c, nil, err.Error())
	}

	return utils.Response200(c, nil, "")
}

func DisabledCategory(c echo.Context) error {
	return c.JSON(http.StatusOK, "Disabled category")
}

func EnabledCategory(c echo.Context) error {
	return c.JSON(http.StatusOK, "Enabled category")
}
