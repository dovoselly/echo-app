package service

import (
	"echo-app/dao"
	"echo-app/model"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateCategory(body model.CategoryCreateBody) error {
	// category BSON
	body.Status = "ENABLE"
	category := model.CategoryBSON{
		ID:          primitive.NewObjectID(),
		Name:        body.Name,
		Description: body.Description,
		Status:      body.Status,
		CreatedAt:   time.Now(),
	}

	// create category
	if err := dao.CreateCategory(category); err != nil {
		return errors.New("can not create new category")
	}

	return nil
}

func GetListCategory() ([]model.CategoryResponse, error) {

	listCategory := make([]model.CategoryResponse, 0)

	// get list category bson
	categoriesBSON, err := dao.GetListCategory()
	if err != nil {
		return listCategory, err
	}

	for _, categoryBSON := range categoriesBSON {
		categoryJSON := model.CategoryResponse{
			ID:          categoryBSON.ID,
			Name:        categoryBSON.Name,
			Description: categoryBSON.Description,
			Status:      categoryBSON.Status,
		}
		listCategory = append(listCategory, categoryJSON)
	}

	return listCategory, nil

}

func GetCategoryByID(ID string) (model.CategoryResponse, error) {
	var (
		category model.CategoryResponse
	)

	// to objectID
	objID, _ := primitive.ObjectIDFromHex(ID)

	// get category by id
	categoryBSON, err := dao.GetCategoryByID(objID)

	category = model.CategoryResponse{
		ID:          categoryBSON.ID,
		Name:        categoryBSON.Name,
		Description: categoryBSON.Description,
		Status:      categoryBSON.Status,
		CreatedAt:   categoryBSON.CreatedAt,
		UpdatedAt:   categoryBSON.UpdatedAt,
	}

	if err != nil {
		return category, err
	}

	return category, nil
}

func UpdateCategoryByID(ID string, body model.CategoryUpdateBody) error {
	objID, _ := primitive.ObjectIDFromHex(ID)

	err := dao.UpdateCategoryByID(objID, body)
	if err != nil {
		return err
	}
	return nil

}

func DeleteCategoryByID(ID string) error {
	// convert id string to objectID
	objID, _ := primitive.ObjectIDFromHex(ID)

	err := dao.DeleteCategoryByID(objID)
	if err != nil {
		return err
	}

	// success
	return nil
}
