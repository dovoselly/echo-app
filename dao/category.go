package dao

import (
	"context"
	"echo-app/database"
	"echo-app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateCategory(category models.Category) error {
	var (
		categoryCol = database.CategoryCol()
		ctx         = context.Background()
	)

	// InsertOne
	_, err := categoryCol.InsertOne(ctx, category)

	return err
}

//func GetListCategory(page, limit int) ([]models.Category, error) {
//	count,_ :=
//}

func GetCategoryByID(ID primitive.ObjectID) (models.Category, error) {
	var (
		categoryCol = database.CategoryCol()
		ctx         = context.Background()
		category    models.Category
	)

	// find category
	filter := bson.M{"_id": ID}
	if err := categoryCol.FindOne(ctx, filter).Decode(&category); err != nil {
		return category, err
	}

	return category, nil
}

func UpdateCategory(ID primitive.ObjectID, body models.CategoryUpdateBody) error {
	var (
		categoryCol = database.CategoryCol()
		ctx         = context.Background()
	)

	update := bson.M{"name": body.Name, "description": body.Description}

	// UpdateOne
	_, err := categoryCol.UpdateOne(ctx, bson.M{"_id": ID}, bson.M{"$set": update})

	if err != nil {
		return err
	}

	return nil
}
