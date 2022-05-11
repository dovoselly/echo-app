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
