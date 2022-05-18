package service

import (
	"echo-app/model"
	"echo-app/util"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct{}

func (u User) ChangePassword(id string, body model.UserChangePassword) error {
	objId, _ := primitive.ObjectIDFromHex(id)
	// check currentPassword
	userBSON, _ := userDAO.GetById(objId)
	if u.checkPasswordHash(body.CurrentPassword, userBSON.Password) != nil {
		return errors.New(utils.CURRENT_PASSWORD_INCORRECT)
	}

	// HashPassword truoc khi update
	newPassword, _ := util.HashPassword(body.NewPassword)

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

	// convert to userInfo
	info = model.UserInfo{
		ID:          user.ID,
		FullName:    user.FullName,
		Email:       user.Email,
		Username:    user.Username,
		Avatar:      user.Avatar,
		Gender:      user.Gender,
		DateOfBirth: user.DateOfBirth,
		Phone:       user.Phone,
		Address:     user.Address,
	}

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
