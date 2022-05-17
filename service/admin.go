package service

import (
	"echo-app/dao"
	"echo-app/model"
	"echo-app/utils"
	"errors"
	"fmt"
)

func AdminLogin(loginBody model.AdminLoginBody) (string, error) {
	// find admin in db
	admin, err := dao.AdminFindByUsername(loginBody.Username)

	if err != nil {
		return "", err
	}

	// verify admin password
	if admin.HashedPassword != loginBody.Password {
		return "", errors.New("wrong password")
	}

	data := map[string]interface{}{
		"id":      admin.ID,
		"isAdmin": true,
	}

	// return JWT token
	token, err := utils.GenerateToken(data)
	if err != nil {
		fmt.Printf(err.Error())
	}
	return token, err
}

func MyProfileAdmin(ID string) (model.Admin, error) {
	doc, err := dao.AdminProfileFindByID(ID)
	if err != nil {
		return doc, err
	}
	return doc, nil
}

func UpdateMyProfileAdmin(ID string, newProfile model.Admin) error {
	err := dao.UpdateMyProfileAdmin(ID, newProfile)
	return err
}
