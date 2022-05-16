package dao

import (
	"echo-app/database"
	"echo-app/models"
	"echo-app/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateReply(insertData models.Reply) error {
	_, err := database.ReplyCol().InsertOne(utils.Ctx, insertData)
	return err
}

func UpdateReply(filter bson.M, updateData bson.M) (*mongo.UpdateResult, error) {
	results, err := database.ReplyCol().UpdateOne(utils.Ctx, filter, updateData)
	return results, err
}

func DeleteReply(filter bson.M) (*mongo.DeleteResult, error) {
	results, err := database.ReplyCol().DeleteOne(utils.Ctx, filter)
	return results, err
}
