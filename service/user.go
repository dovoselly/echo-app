package service

import (
	"echo-app/model"
	"echo-app/util"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/pkg/errors"
)

type User struct{}

func (u User) ChangePassword(ID string, body model.UserChangePassword) error {
	// check currentPassword
	objId, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}

	userBSON, _ := userDAO.GetById(objId)
	if util.CheckPasswordHash(body.CurrentPassword, userBSON.Password) != nil {
		return errors.New("CurrentPassword is incorrect")
	}

	// HashPassword truoc khi update
	newPassword, _ := util.HashPassword(body.NewPassword)

	// update password
	err = userDAO.UpdatePassword(objId, newPassword)

	if err != nil {
		return err
	}

	return nil
}

func (u User) GetInfo(ID string) (model.UserInfo, error) {
	var (
		info model.UserInfo
	)

	objId, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return model.UserInfo{}, err
	}

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

func (u User) UpdateInfo(ID string, body model.UserUpdate) error {
	objId, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}

	bodyBSON := model.UserInfoBSON{
		FullName:    body.FullName,
		Email:       body.Email,
		Phone:       body.Phone,
		DateOfBirth: body.DateOfBirth,
		Gender:      body.Gender,
		Address:     body.Address,
	}

	// update info
	if err := userDAO.UpdateInfo(objId, bodyBSON); err != nil {
		return err
	}

	return nil
}
