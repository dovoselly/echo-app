package dao

import (
	"echo-app/database"
	"echo-app/model"
	"echo-app/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateReply(insertData model.Reply) error {
	_, err := database.ReplyCol().InsertOne(util.CONTEXT, insertData)
	return err
}

func UpdateReply(filter bson.M, updateData bson.M) (*mongo.UpdateResult, error) {
	results, err := database.ReplyCol().UpdateOne(util.CONTEXT, filter, updateData)
	return results, err
}

func DeleteReply(filter bson.M) (*mongo.DeleteResult, error) {
	results, err := database.ReplyCol().DeleteOne(util.CONTEXT, filter)
	return results, err
}
