package dao

import (
	"context"
	"echo-app/database"
	"echo-app/models"
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
