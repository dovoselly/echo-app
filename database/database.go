package database

import (
	"context"
	"echo-app/config"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var database *mongo.Database
var userColName = "users"
var adminColName = "admin"

func Connect() {
	var env = config.GetEnv()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(env.Database.URI))
	if err != nil {
		fmt.Println(err.Error())
	}

	database = client.Database(env.Database.Name)
}
