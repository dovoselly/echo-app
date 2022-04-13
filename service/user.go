package service

import (
	"context"
	"echo-app/model"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
)

type UserService interface {
	CreateUser()
	GetAllUsers()
}

type Response struct {
	Message string `json:"message"`
}

var collection *mongo.Collection

func runDB() interface{} {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://vothanhdo2602:MF0No0c1MrvJT384@cluster0.mggsp.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		fmt.Println(err)
	}
	collection = client.Database("myFirstDatabase").Collection("users")
	return ctx
}

func checkCollectionInstance() {
	if collection == nil {
		runDB()
	}
}

func CreateUser(ctxEcho echo.Context) error {
	checkCollectionInstance()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	payload := new(model.User)
	if err := ctxEcho.Bind(payload); err != nil {
		return ctxEcho.JSON(http.StatusBadRequest, Response{
			Message: "invalid value",
		})
	}
	_, err := collection.InsertOne(ctx, payload)
	if err != nil {
		return ctxEcho.JSON(http.StatusInternalServerError, err)
	}
	return ctxEcho.JSON(http.StatusOK, Response{
		Message: "Create successfully",
	})
}

func AllUsers(ctxEcho echo.Context) error {
	checkCollectionInstance()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return ctxEcho.JSON(http.StatusInternalServerError, err)
	}
	var people []model.User
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var person model.User
		cursor.Decode(&person)
		people = append(people, person)
	}
	if err := cursor.Err(); err != nil {
		return ctxEcho.JSON(http.StatusInternalServerError, err)
	}
	return ctxEcho.JSON(http.StatusOK, people)
}

func GetUserById(ctxEcho echo.Context) error {
	checkCollectionInstance()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user model.User
	id, _ := primitive.ObjectIDFromHex(ctxEcho.Param("id"))
	err := collection.FindOne(ctx, model.User{ID: id}).Decode(&user)
	if err != nil {
		ctxEcho.JSON(http.StatusInternalServerError, err)
	}
	return ctxEcho.JSON(http.StatusOK, user)
}

func UpdateUserById(ctxEcho echo.Context) error {
	checkCollectionInstance()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	payload := new(model.User)
	if err := ctxEcho.Bind(payload); err != nil {
		return ctxEcho.JSON(http.StatusBadRequest, Response{
			Message: "invalid value",
		})
	}
	id, _ := primitive.ObjectIDFromHex(ctxEcho.Param("id"))
	result, err := collection.UpdateOne(ctx, model.User{ID: id}, bson.M{"$set": payload})
	if err != nil {
		ctxEcho.JSON(http.StatusInternalServerError, err)
	}
	return ctxEcho.JSON(http.StatusOK, result)
}

func DeleteUserById(ctxEcho echo.Context) error {
	checkCollectionInstance()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	id, _ := primitive.ObjectIDFromHex(ctxEcho.Param("id"))
	result, err := collection.DeleteOne(ctx, model.User{ID: id})
	if err != nil {
		ctxEcho.JSON(http.StatusInternalServerError, err)
	}
	return ctxEcho.JSON(http.StatusOK, result)
}
