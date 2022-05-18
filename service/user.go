package service

import (
	"echo-app/model"
	"echo-app/utils"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct{}

func (u User) ChangePassword(id string, body model.UserChangePassword) error {
	objId, _ := primitive.ObjectIDFromHex(id)
	// check currentPassword
	userBSON, _ := userDAO.GetById(objId)
	if u.checkPasswordHash(body.CurrentPassword, userBSON.Password) != nil {
		return errors.New(utils.CURRENT_PASSWORD_INCORRECT)
	}

	// hash password before update
	newPassword, _ := u.hashPassword(body.NewPassword)

	// update password
	err := userDAO.UpdatePassword(objId, newPassword)

	if err != nil {
		return err
	}

	return nil
}

func (u User) GetInfo(id string) (model.UserInfo, error) {
	var (
		info model.UserInfo
	)

	objId, _ := primitive.ObjectIDFromHex(id)

	// get user
	user, err := userDAO.GetInfo(objId)
	if err != nil {
		return info, err
	}

	// convert userBson to userInfo
	info = user.ConvertToJSON()

	return info, nil
}

func (u User) UpdateInfo(id string, body model.UserUpdate) error {
	objId, _ := primitive.ObjectIDFromHex(id)
	// convert userUpdate to userBson
	bodyBSON := body.ConvertToBSON()

	// update info
	if err := userDAO.UpdateInfo(objId, bodyBSON); err != nil {
		return err
	}

	return nil
}

func (u User) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (u User) checkPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}
