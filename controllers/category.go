package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateCategory(c echo.Context) error {
	return c.JSON(http.StatusOK, "Create new category")
}

func GetListCategory(c echo.Context) error {
	return c.JSON(http.StatusOK, "Get all category")
}

func GetCategoryByID(c echo.Context) error {
	return c.JSON(http.StatusOK, "Get category detail")
}

func UpdateCategory(c echo.Context) error {
	return c.JSON(http.StatusOK, "Update category")
}

func DisabledCategory(c echo.Context) error {
	return c.JSON(http.StatusOK, "Disabled category")
}

func EnabledCategory(c echo.Context) error {
	return c.JSON(http.StatusOK, "Enabled category")
}
