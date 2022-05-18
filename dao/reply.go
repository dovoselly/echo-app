package dao

import (
	"echo-app/database"
	"echo-app/model"
	"echo-app/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Reply struct{}

func (Reply) CreateReply(insertData model.Reply) (string, error) {
	result, err := database.ReplyCol().InsertOne(utils.Ctx, insertData)
	return result.InsertedID.(string), err
}

func (Reply) UpdateReply(userOjbID primitive.ObjectID, replyOjbID primitive.ObjectID, updateData model.Reply) (*mongo.UpdateResult, error) {
	var (
		filter        = bson.M{"_id": replyOjbID, "userId": userOjbID}
		updateDataSet = bson.M{"$set": updateData}
	)

	results, err := database.ReplyCol().UpdateOne(utils.Ctx, filter, updateDataSet)
	return results, err
}

func (Reply) DeleteReply(userOjbID primitive.ObjectID, replyOjbID primitive.ObjectID) (*mongo.DeleteResult, error) {
	var filter = bson.M{"_id": replyOjbID, "userId": userOjbID}

	results, err := database.ReplyCol().DeleteOne(utils.Ctx, filter)
	return results, err
}
