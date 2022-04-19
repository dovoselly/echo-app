package database

import (
	"context"
	"echo-app/config"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var database *mongo.Database
var userColName = "users"
var adminColName = "admin"

func RunDB() {
	var env = config.GetEnv()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(env.Database.URI))
	if err != nil {
		fmt.Println(err.Error())
	}

	database = client.Database(env.Database.Name)
}

func GetUserCol() *mongo.Collection {
	return database.Collection(userColName)
}

func GetAdminCol() *mongo.Collection {
	return database.Collection(adminColName)
}
