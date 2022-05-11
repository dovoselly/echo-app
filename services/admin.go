package services

import (
	"echo-app/dao"
	"echo-app/middlewares"
	"echo-app/models"
	"errors"
)

func AdminLogin(loginBody models.AdminLoginBody) (string, error) {
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
	return middlewares.GenerateToken(data)
}

//func MyProfile(ID string) (models.Admin, error) {
//
//}
