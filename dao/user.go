package dao

import (
	"context"
	"echo-app/database"
	"echo-app/models"

	"go.mongodb.org/mongo-driver/bson"
)

func UserRegister(doc models.UserBSON) (models.UserBSON, error) {
	var (
		userCol = database.UserCol()
		ctx     = context.Background()
	)

	// Insert one
	_, err := userCol.InsertOne(ctx, doc)
	return doc, err
}

func GetUserByUsername(username string) (models.UserBSON, error) {
	var (
		userCol = database.UserCol()
		ctx     = context.Background()
		user    models.UserBSON
	)

	// filter
	filter := bson.M{"username": username}

	// FindOne
	err := userCol.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		return user, err
	}

	return user, nil
}
