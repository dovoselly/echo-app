package services

import (
	"context"
	"echo-app/dao"
	"echo-app/database"
	"echo-app/models"
	"echo-app/utils"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func UserRegister(payload models.UserRegister) (models.UserBSON, error) {
	var (
		collection = database.UserCol()
		ctx        = context.Background()
		user       *models.UserRepsonse
	)

	// check exist email and username
	errExist := collection.FindOne(ctx, bson.M{"$or": []bson.M{{"email": payload.Email}, {"username": payload.Username}}}).Decode(&user)
	fmt.Println("error: ", errExist)

	if user != nil {
		if user.Email == payload.Email && user.Username == payload.Username {
			return models.UserBSON{}, errors.New("email and username is already")
		}
		if user.Email == payload.Email {
			return models.UserBSON{}, errors.New("email is already")
		}
		if user.Username == payload.Username {
			return models.UserBSON{}, errors.New("username is already")
		}
	}

	// HashPassword
	payload.Password, _ = utils.HashPassword(payload.Password)

	// default status
	payload.Status = "ACTIVE"

	//Create user
	doc, err := dao.UserRegister(payload.ConvertToBSON())
	if err != nil {
		err = errors.New("khong the tao user")
		return doc, err
	}

	return doc, err

}

func Login(user models.UserLogin) (string, error) {

	// FInd user by username
	userBSON, err := dao.GetUserByUsername(user.Username)
	if err != nil {
		return "", errors.New("Email not existed in db")
	}

	// verify user password
	if utils.CheckPasswordHash(user.Password, userBSON.Password) != nil {
		return "", errors.New("Wrong password")
	}

	// JWT payload
	data := map[string]interface{}{
		"id": userBSON.ID,
	}

	// Genderate user token
	token, err := utils.GenerateUserToken(data)
	if err != nil {
		return "", errors.New("GenerateUserToken failed")
	}

	return token, nil

}
