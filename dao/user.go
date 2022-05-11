package dao

import (
	"context"
	"echo-app/database"
	"echo-app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func GetUserById(ID primitive.ObjectID) (models.UserBSON, error) {
	var (
		userCol = database.UserCol()
		ctx     = context.Background()
		user    models.UserBSON
	)

	err := userCol.FindOne(ctx, bson.M{"_id": ID}).Decode(&user)
	if err != nil {
		return models.UserBSON{}, err
	}
	return user, nil

}

func UpdateUserPassword(ID primitive.ObjectID, newPassword string) error {
	var (
		userCol = database.UserCol()
		ctx     = context.Background()
	)

	_, err := userCol.UpdateOne(
		ctx,
		bson.M{"_id": ID},
		bson.D{
			{"$set", bson.D{{"password", newPassword}}},
		},
	)

	if err != nil {
		return err
	}
	return nil
}
