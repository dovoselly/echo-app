package dao

import (
	"context"
	"echo-app/database"
	"echo-app/models"
	"fmt"
)

func UserRegister(doc models.UserBSON) (models.UserBSON, error) {
	fmt.Println("dao")
	var (
		collection = database.UserCol()
		ctx        = context.Background()
	)

	// Insert one
	_, err := collection.InsertOne(ctx, doc)
	return doc, err
}
