package controllers

import (
	"echo-app/models"
	"echo-app/services"
	"echo-app/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateCategory(c echo.Context) error {
	var body = c.Get("body").(models.CategoryCreateBody)

	// process data
	err := services.CreateCategory(body)

	// if err
	if err != nil {
		return utils.Response400(c, nil, err.Error())
	}

	return utils.Response200(c, body, "")
}

func GetListCategory(c echo.Context) error {
	categories, err := services.GetListCategory()
	if err != nil {
		return utils.Response400(c, nil, err.Error())
	}
	return utils.Response200(c, categories, "")
}

func GetCategoryByID(c echo.Context) error {
	var strID = c.Get("id").(string)

	// process
	category, err := services.GetCategoryByID(strID)

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
	err := services.UpdateCategoryByID(ID, body)
	if err != nil {
		return utils.Response400(c, nil, err.Error())
	}

	return utils.Response200(c, nil, "")

}

func DeleteCategoryByID(c echo.Context) error {
	var id = c.Get("id").(string)

	//process
	err := services.DeleteCategoryByID(id)
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
