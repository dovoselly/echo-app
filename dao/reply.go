package dao

import (
	"echo-app/database"
	"echo-app/models"
	"echo-app/utils"
)

func CreateReply(insertData models.Reply) error {
	_, err := database.ReplyCol().InsertOne(utils.Ctx, insertData)
	return err
}
