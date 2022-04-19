package service

import (
	"context"
	"echo-app/database"
	"echo-app/model"
	"echo-app/util"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

var ctx = context.TODO()

func CreateUser(insertData model.User) string {
	_, err := database.GetUserCol().InsertOne(ctx, insertData)
	if err != nil {
		return err.Error()
	}
	return "Create successfully"
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

func AllUsers(c echo.Context, page int, limit int) []model.User {
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

func GetUserById(c echo.Context, id primitive.ObjectID) error {
	var user model.User
	err := database.GetUserCol().FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return c.JSON(http.StatusOK, util.Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, user)
}

func UpdateUserById(c echo.Context, id primitive.ObjectID, insertData model.User) error {
	result, err := database.GetUserCol().UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": insertData})
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteUserById(c echo.Context, id primitive.ObjectID) error {
	result, err := database.GetUserCol().DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}
