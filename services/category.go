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

	category := models.Category{
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

//func GetAllCategory()

func GetCategoryByID(ID string) (models.Category, error) {
	// to objectID
	objID, _ := primitive.ObjectIDFromHex(ID)

	// get category by id
	category, err := dao.GetCategoryByID(objID)

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
