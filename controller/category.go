package controller

import (
	"echo-app/model"
	"echo-app/util"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct{}

func (ca Category) Create(c echo.Context) error {
	var body = c.Get("body").(model.CategoryCreateBody)

	// process data
	categoryID, err := categoryService.Create(body)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, categoryID, "")
}

func (ca Category) GetList(c echo.Context) error {
	categories, err := categoryService.GetList()
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, categories, "")
}

func (ca Category) GetByID(c echo.Context) error {
	var (
		id = c.Get("id").(string)
	)
	// to objectID
	objID, _ := primitive.ObjectIDFromHex(id)

	// process
	category, err := categoryService.GetByID(objID)
	// if error
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, category, "")
}

func (ca Category) UpdateByID(c echo.Context) error {
	var (
		id   = c.Get("id").(string)
		body = c.Get("body").(model.CategoryUpdateBody)
	)

	objID, _ := primitive.ObjectIDFromHex(id)

	// process data
	categoryID, err := categoryService.UpdateByID(objID, body)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, bson.M{"_id": categoryID}, "")
}

func (ca Category) DeleteByID(c echo.Context) error {
	var (
		id = c.Get("id").(string)
	)

	objID, _ := primitive.ObjectIDFromHex(id)

	//process
	err := categoryService.DeleteByID(objID)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, nil, "")
}

func (ca Category) UpdateStatus(c echo.Context) error {
	var (
		id = c.Get("id").(string)
	)

	objID, _ := primitive.ObjectIDFromHex(id)

	//process
	categoryID, err := categoryService.UpdateStatus(objID)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, bson.M{"_id": categoryID}, "")
}
