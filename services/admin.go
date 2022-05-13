package services

import (
	"echo-app/dao"
	"echo-app/models"
	"echo-app/utils"
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
	return utils.GenerateToken(data), err
}

func MyProfileAdmin(ID string) (models.Admin, error) {
	doc, err := dao.AdminProfileFindByID(ID)
	if err != nil {
		return doc, err
	}
	return doc, nil
}

//func GetAdminProfileByID(id string) (models.Admin, error) {
//	// to objectID
//	objID, _ := primitive.ObjectIDFromHex(id)
//
//	admin, err := dao.GetAdminProfileFindByID(objID)
//	if err != nil {
//		return admin, err
//	}
//	return admin, nil
//}