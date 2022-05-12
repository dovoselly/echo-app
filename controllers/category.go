package controllers

import (
	"echo-app/models"
	"echo-app/services"
	"echo-app/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateCategory(c echo.Context) error {
	var body = c.Get("categoryBody").(models.CategoryCreateBody)

	// process data
	err := services.CreateCategory(body)

	// if err
	if err != nil {
		return utils.Response400(c, nil, err.Error())
	}

	return utils.Response200(c, body, "")
}

func GetListCategory(c echo.Context) error {
	return c.JSON(http.StatusOK, "Get all category")
}

func GetCategoryByID(c echo.Context) error {
	var strID = c.Get("strID").(string)

	// process
	category, err := services.GetCategoryByID(strID)

	// if error
	if err != nil {
		return utils.Response400(c, nil, err.Error())
	}

	return utils.Response200(c, category, "")
}

func UpdateCategory(c echo.Context) error {
	var (
		ID   = c.Get("id").(string)
		body = c.Get("body").(models.CategoryUpdateBody)
	)

	// process data
	err := services.UpdateCategory(ID, body)
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
