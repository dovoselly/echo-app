package service

import (
	"echo-app/model"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Category struct{}

func (c Category) CreateCategory(body model.CategoryCreateBody) error {
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
	if err := categoryDao.CreateCategory(category); err != nil {
		return errors.New("can not create new category")
	}

	return nil
}

func (c Category) GetListCategory() ([]model.CategoryResponse, error) {

	listCategory := make([]model.CategoryResponse, 0)

	// get list category bson
	categoriesBSON, err := categoryDao.GetListCategory()
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

func (c Category) GetCategoryByID(ID string) (model.CategoryResponse, error) {
	var (
		category model.CategoryResponse
	)

	// to objectID
	objID, _ := primitive.ObjectIDFromHex(ID)

	// get category by id
	categoryBSON, err := categoryDao.GetCategoryByID(objID)

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

func (c Category) UpdateCategoryByID(ID string, body model.CategoryUpdateBody) error {
	objID, _ := primitive.ObjectIDFromHex(ID)

	err := categoryDao.UpdateCategoryByID(objID, body)
	if err != nil {
		return err
	}
	return nil

}

func (c Category) DeleteCategoryByID(ID string) error {
	// convert id string to objectID
	objID, _ := primitive.ObjectIDFromHex(ID)

	err := categoryDao.DeleteCategoryByID(objID)
	if err != nil {
		return err
	}

	// success
	return nil
}
