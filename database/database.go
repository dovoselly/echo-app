package database

import (
	"context"
	"echo-app/util"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

var database *mongo.Database

func RunDB() {
	util.Dotenv()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("DB_CONNECT_URI")))
	if err != nil {
		fmt.Println(err.Error())
	}

	database = client.Database("myFirstDatabase")
}

func GetUserCol() *mongo.Collection {
	return database.Collection("users")
}
