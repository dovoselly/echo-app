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

func CreateUser(c echo.Context, insertData *model.User, isAdmin bool) error {
	if isAdmin {
		insertData.Role = "admin"
	} else {
		insertData.Role = "user"
	}
	_, err := database.GetUserCol().InsertOne(ctx, insertData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, util.Response{
		Message: "Create successfully",
	})
}

func GetUserByEmail(c echo.Context, email string) (model.User, error) {
	var user model.User
	err := database.GetUserCol().FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return user, c.JSON(http.StatusBadRequest, err.Error())
	}
	return user, nil
}

func AllUsers(c echo.Context, page int64, limit int64) error {
	// options query
	optionsQuery := new(options.FindOptions)
	optionsQuery.SetSkip((page - 1) * limit)
	optionsQuery.SetLimit(limit)

	// Find users
	cursor, err := database.GetUserCol().Find(ctx, bson.M{}, optionsQuery)
	if err != nil {
		fmt.Println(err.Error(), "111111111")
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Decode found documents
	var people []model.User
	if err := cursor.All(ctx, &people); err != nil {
		fmt.Println(err.Error(), "2222222222222")
		return c.JSON(http.StatusBadRequest, people)
	}

	return c.JSON(http.StatusOK, people)
}

func GetUserById(c echo.Context, id primitive.ObjectID) error {
	var user model.User
	err := database.GetUserCol().FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func UpdateUserById(c echo.Context, id primitive.ObjectID, insertData model.User) error {
	result, err := database.GetUserCol().UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": insertData})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteUserById(c echo.Context, id primitive.ObjectID) error {
	result, err := database.GetUserCol().DeleteOne(ctx, bson.M{"set": id})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
