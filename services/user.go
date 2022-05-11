package services

import (
	"echo-app/dao"
	"echo-app/models"
	"echo-app/utils"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UserChangePassword(ID string, body models.UserChangePassword) error {

	// convert id string to primitive.ObjectID
	id, _ := primitive.ObjectIDFromHex(ID)

	// check currentPassword
	userBSON, _ := dao.GetUserById(id)
	if utils.CheckPasswordHash(body.CurrentPassword, userBSON.Password) != nil {
		return errors.New("CurrentPassword is incorrect")
	}

	// HashPassword truoc khi update
	newPassword, _ := utils.HashPassword(body.NewPassword)

	// update password
	err := dao.UpdateUserPassword(id, newPassword)

	if err != nil {
		return err
	}

	return nil
}
