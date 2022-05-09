package database

import "go.mongodb.org/mongo-driver/mongo"

func UserCol() *mongo.Collection {
	return database.Collection(userColName)
}

func AdminCol() *mongo.Collection {
	return database.Collection(adminColName)
}
