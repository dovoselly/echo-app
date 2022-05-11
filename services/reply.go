package services

import (
	"echo-app/dao"
	"echo-app/models"
	"echo-app/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateReply(id primitive.ObjectID, payload models.CreateReply) error {
	//init insert data
	insertData := models.Reply{
		Content:   payload.Content,
		ReviewId:  id,
		CreatedAt: utils.CurrentDateTime(),
		UpdatedAt: utils.CurrentDateTime(),
	}
	err := dao.CreateReply(insertData)
	return err
}
