package service

import (
	"echo-app/model"
	"echo-app/utils"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct{}

func (u User) ChangePassword(id primitive.ObjectID, body model.UserChangePassword) error {
	// check currentPassword
	userBSON, _ := userDAO.GetById(id)
	if u.checkPasswordHash(body.CurrentPassword, userBSON.Password) != nil {
		return errors.New(utils.CURRENT_PASSWORD_INCORRECT)
	}

	// hash password before update
	newPassword, _ := u.hashPassword(body.NewPassword)

	// update password
	err := userDAO.UpdatePassword(id, newPassword)

	if err != nil {
		return err
	}

	return nil
}

func (u User) GetInfo(id primitive.ObjectID) (model.UserInfo, error) {
	var (
		info model.UserInfo
	)

	// get user
	user, err := userDAO.GetInfo(id)
	if err != nil {
		return info, err
	}

	// convert userBson to userInfo
	info = user.ConvertToJSON()

	return info, nil
}

func (u User) UpdateInfo(id primitive.ObjectID, body model.UserUpdate) error {
	// convert userUpdate to userBson
	bodyBSON := body.ConvertToBSON()

	// update info
	if err := userDAO.UpdateInfo(id, bodyBSON); err != nil {
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
