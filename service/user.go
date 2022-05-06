package service

import (
	"context"
	"echo-app/database"
	"echo-app/model"
	"echo-app/util"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()

func CreateUser(insertData model.User) string {
	_, err := database.GetUserCol().InsertOne(ctx, insertData)
	if err != nil {
		return err.Error()
	}
	return util.CreateSuccessFully
}

func GetUserByEmail(email string) model.User {
	var user model.User
	err := database.GetUserCol().FindOne(ctx, bson.M{"email": email}).Decode(&user)
	fmt.Println(user)
	if err != nil {
		fmt.Println(err.Error())
	}
	return user
}

func AllUsers(page int, limit int) []model.User {
	var allUsers []model.User

	// options query
	optionsQuery := new(options.FindOptions)
	optionsQuery.SetSkip(int64(page * limit))
	optionsQuery.SetLimit(int64(limit))

	// Find users
	cursor, err := database.GetUserCol().Find(ctx, bson.M{}, optionsQuery)
	if err != nil {
		fmt.Println(err.Error())
		return allUsers
	}

	// Decode found documents
	if err := cursor.All(ctx, &allUsers); err != nil {
		fmt.Println(err.Error())
		return allUsers
	}

	return allUsers
}

func GetUserById(id primitive.ObjectID) model.User {
	var user model.User
	err := database.GetUserCol().FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		fmt.Println(err.Error())
		return user
	}
	return user
}

func UpdateUserById(id primitive.ObjectID, insertData model.UserUpdate) string {
	fmt.Println(id, insertData)
	result, err := database.GetUserCol().UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": insertData})
	fmt.Println(*result)
	if err != nil {
		return err.Error()
	}
	if result.MatchedCount == 0 {
		return util.UserNotFound
	}
	return util.UpdateSuccessFully
}

func DeleteUserById(id primitive.ObjectID) string {
	result, err := database.GetUserCol().DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err.Error()
	}
	if result.DeletedCount == 0 {
		return util.UserNotFound
	}
	return util.DeleteSuccessFully
}
