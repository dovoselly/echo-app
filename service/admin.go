package service

import (
	"echo-app/model"
	"echo-app/utils"
	"errors"
	"fmt"
)

type Admin struct{}

func (a Admin) AdminLogin(body model.AdminLoginBody) (string, error) {
	// find admin in db
	admin, err := adminDao.FindByUsername(body.Username)

	if err != nil {
		return "", err
	}

	// verify admin password
	if admin.HashedPassword != body.Password {
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

func (a Admin) GetMyProfile(ID string) (model.Admin, error) {
	doc, err := adminDao.ProfileFindByID(ID)
	if err != nil {
		return doc, err
	}
	return doc, nil
}

func (a Admin) UpdateMyProfile(ID string, newProfile model.Admin) error {
	err := adminDao.UpdateMyProfile(ID, newProfile)
	return err
}
