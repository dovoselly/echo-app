package dao

import (
	"context"
	"echo-app/database"
	"echo-app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateReply(insertData models.Reply) error {
	var (
		ctx = context.Background()
	)
	_, err := database.ReplyCol().InsertOne(ctx, insertData)
	return err
}

func UpdateReply(filter bson.M, updateData bson.M) (*mongo.UpdateResult, error) {
	var (
		ctx = context.Background()
	)

	results, err := database.ReplyCol().UpdateOne(ctx, filter, updateData)
	return results, err
}

func DeleteReply(filter bson.M) (*mongo.DeleteResult, error) {
	var (
		ctx = context.Background()
	)

	results, err := database.ReplyCol().DeleteOne(ctx, filter)
	return results, err
}
