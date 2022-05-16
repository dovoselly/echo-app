package services

import (
	"echo-app/dao"
	"echo-app/models"
	"echo-app/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateReply(userId primitive.ObjectID, reviewId primitive.ObjectID, body models.CreateReply) error {
	//init insert data
	insertData := models.Reply{
		UserId:    userId,
		Content:   body.Content,
		ReviewId:  reviewId,
		CreatedAt: utils.CurrentDateTime(),
		UpdatedAt: utils.CurrentDateTime(),
	}
	err := dao.CreateReply(insertData)
	return err
}

func UpdateReply(userId primitive.ObjectID, replyId primitive.ObjectID, body models.CreateReply) (*mongo.UpdateResult, error) {
	filter := bson.M{"_id": replyId, "userId": userId}

	updateData := bson.M{"$set": models.Reply{
		Content:   body.Content,
		UpdatedAt: utils.CurrentDateTime(),
	}}

	results, err := dao.UpdateReply(filter, updateData)
	return results, err
}

func DeleteReply(userId primitive.ObjectID, replyId primitive.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.M{"_id": replyId, "userId": userId}

	results, err := dao.DeleteReply(filter)
	return results, err
}
