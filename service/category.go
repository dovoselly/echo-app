package service

import (
	"echo-app/dao"
	"echo-app/models"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateCategory(body models.CategoryCreateBody) error {
	// category BSON
	body.Status = "ENABLE"
	category := models.CategoryBSON{
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

func GetListCategory() ([]models.CategoryResponse, error) {

	listCategory := make([]models.CategoryResponse, 0)

	// get list category bson
	categoriesBSON, err := dao.GetListCategory()
	if err != nil {
		return listCategory, err
	}

	for _, categoryBSON := range categoriesBSON {
		categoryJSON := models.CategoryResponse{
			ID:          categoryBSON.ID,
			Name:        categoryBSON.Name,
			Description: categoryBSON.Description,
			Status:      categoryBSON.Status,
		}
		listCategory = append(listCategory, categoryJSON)
	}

	return listCategory, nil

}

func GetCategoryByID(ID string) (models.CategoryResponse, error) {
	var (
		category models.CategoryResponse
	)

	// to objectID
	objID, _ := primitive.ObjectIDFromHex(ID)

	// get category by id
	categoryBSON, err := dao.GetCategoryByID(objID)

	category = models.CategoryResponse{
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

func UpdateCategoryByID(ID string, body models.CategoryUpdateBody) error {
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
