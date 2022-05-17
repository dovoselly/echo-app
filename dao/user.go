package dao

import (
	"context"
	"echo-app/database"
	"echo-app/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UserRegister(doc model.UserBSON) (model.UserBSON, error) {
	var (
		userCol = database.UserCol()
		ctx     = context.Background()
	)

	// Insert one
	_, err := userCol.InsertOne(ctx, doc)
	return doc, err
}

func GetUserByUsername(username string) (model.UserBSON, error) {
	var (
		userCol = database.UserCol()
		ctx     = context.Background()
		user    model.UserBSON
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

func GetUserById(ID primitive.ObjectID) (model.UserBSON, error) {
	var (
		userCol = database.UserCol()
		ctx     = context.Background()
		user    model.UserBSON
	)

	err := userCol.FindOne(ctx, bson.M{"_id": ID}).Decode(&user)
	if err != nil {
		return model.UserBSON{}, err
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

func GetInfoUser(ID primitive.ObjectID) (model.UserBSON, error) {
	var (
		userCol = database.UserCol()
		ctx     = context.Background()
		user    model.UserBSON
	)
	filter := bson.M{"_id": ID}
	err := userCol.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func UpdateInfoUser(ID primitive.ObjectID, body model.UserInfoBSON) error {
	var (
		userCol = database.UserCol()
		ctx     = context.Background()
	)

	filter := bson.M{"_id": ID}

	_, err := userCol.UpdateOne(ctx, filter, bson.M{"$set": body})
	if err != nil {
		return err
	}

	return nil
}
