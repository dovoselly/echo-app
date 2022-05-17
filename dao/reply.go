package dao

import (
	"context"
	"echo-app/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateReply(insertData model.Reply) error {
	_, err := database.ReplyCol().InsertOne(utils.Ctx, insertData)
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
