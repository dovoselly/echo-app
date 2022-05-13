package services

import (
	"echo-app/dao"
	"echo-app/models"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func CreateCategory(categoryBody models.CategoryCreateBody) error {
	// category BSON

	category := models.CategoryBSON{
		ID:          primitive.NewObjectID(),
		Name:        categoryBody.Name,
		Description: categoryBody.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// create bot
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

func UpdateCategory(ID string, body models.CategoryUpdateBody) error {
	objID, _ := primitive.ObjectIDFromHex(ID)

	err := dao.UpdateCategory(objID, body)
	if err != nil {
		return err
	}
	return nil

}
