package controller

import (
	"echo-app/model"
	"echo-app/util"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Category struct{}

func (Category) Create(c echo.Context) error {
	var body = c.Get("body").(model.CategoryCreateBody)

	// process data
	if err := categoryService.CreateCategory(body); err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, body, "")
}

func (Category) GetList(c echo.Context) error {
	categories, err := categoryService.GetListCategory()
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}
	return util.Response200(c, categories, "")

	//if categories, err := service.GetListCategory(); err != nil {
	//	return util.Response400(c, nil, err.Error())
	//}
	//return util.Response200(c, categories, "")

}

func (Category) GetByID(c echo.Context) error {
	var strID = c.Get("id").(string)

	// process
	category, err := categoryService.GetCategoryByID(strID)

	// if error
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, category, "")
}

func (Category) UpdateByID(c echo.Context) error {
	var (
		ID   = c.Get("id").(string)
		body = c.Get("body").(model.CategoryUpdateBody)
	)

	// process data
	err := categoryService.UpdateCategoryByID(ID, body)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, nil, "")

}

func (Category) DeleteByID(c echo.Context) error {
	var id = c.Get("id").(string)

	//process
	err := categoryService.DeleteCategoryByID(id)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, nil, "")
}

func (Category) Disabled(c echo.Context) error {
	return c.JSON(http.StatusOK, "Disabled category")
}

func (Category) Enabled(c echo.Context) error {
	return c.JSON(http.StatusOK, "Enabled category")
}
