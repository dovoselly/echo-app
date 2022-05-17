package service

import (
	"echo-app/dao"
	"echo-app/model"
	"echo-app/utils"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ChangeUserPassword(ID primitive.ObjectID, body model.UserChangePassword) error {
	// check currentPassword
	userBSON, _ := dao.GetUserById(ID)
	if utils.CheckPasswordHash(body.CurrentPassword, userBSON.Password) != nil {
		return errors.New("CurrentPassword is incorrect")
	}

	// HashPassword truoc khi update
	newPassword, _ := utils.HashPassword(body.NewPassword)

	// update password
	err := dao.UpdateUserPassword(ID, newPassword)

	if err != nil {
		return err
	}

	return nil
}

func GetUserInfo(ID primitive.ObjectID) (model.UserInfo, error) {
	var (
		info model.UserInfo
	)

	// get user
	user, err := dao.GetInfoUser(ID)
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

func UpdateUserInfo(ID primitive.ObjectID, body model.UserUpdate) error {

	bodyBSON := model.UserInfoBSON{
		FullName:    body.FullName,
		Email:       body.Email,
		Phone:       body.Phone,
		DateOfBirth: body.DateOfBirth,
		Gender:      body.Gender,
		Address:     body.Address,
	}

	// update info
	if err := dao.UpdateInfoUser(ID, bodyBSON); err != nil {
		return err
	}

	return nil
}
