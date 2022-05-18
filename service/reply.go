package service

import (
	"echo-app/dao"
	"echo-app/model"
	"echo-app/util"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateReply(userId primitive.ObjectID, reviewId primitive.ObjectID, body model.CreateReply) error {
	//init insert data
	insertData := model.Reply{
		UserId:    userId,
		Content:   body.Content,
		ReviewId:  reviewId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := dao.CreateReply(insertData)
	return err
}

func UpdateReply(userId primitive.ObjectID, replyId primitive.ObjectID, body model.CreateReply) (*mongo.UpdateResult, error) {
	filter := bson.M{"_id": replyId, "userId": userId}

	updateData := bson.M{"$set": model.Reply{
		Content:   body.Content,
		UpdatedAt: util.CurrentDateTime(),
	}}

	results, err := dao.UpdateReply(filter, updateData)
	return results, err
}

func DeleteReply(userId primitive.ObjectID, replyId primitive.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.M{"_id": replyId, "userId": userId}

	results, err := dao.DeleteReply(filter)
	return results, err
}
