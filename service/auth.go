package service

import (
	"context"
	"echo-app/database"
	"echo-app/model"
	"echo-app/util"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func UserRegister(payload model.UserRegister) (model.UserBSON, error) {
	var (
		collection = database.UserCol()
		ctx        = context.Background()
		user       *model.UserResponse
	)

	// check exist email and username
	errExist := collection.FindOne(ctx, bson.M{"$or": []bson.M{{"email": payload.Email}, {"username": payload.Username}}}).Decode(&user)

	fmt.Println("error: ", errExist)

	if user != nil {
		if user.Email == payload.Email && user.Username == payload.Username {
			return model.UserBSON{}, errors.New("email and username is already")
		}
		if user.Email == payload.Email {
			return model.UserBSON{}, errors.New("email is already")
		}
		if user.Username == payload.Username {
			return model.UserBSON{}, errors.New("username is already")
		}
	}

	// HashPassword
	payload.Password, _ = util.HashPassword(payload.Password)

	// default status
	payload.Status = "ACTIVE"

	//Create user
	doc, err := userDAO.Register(payload.ConvertToBSON())
	if err != nil {
		err = errors.New("khong the tao user")
		return doc, err
	}

	return doc, err

}

func Login(user model.UserLogin) (string, error) {

	// FInd user by username
	userBSON, err := userDAO.GetByUsername(user.Username)
	if err != nil {
		return "", errors.New("Email not existed in db")
	}

	// verify user password
	if util.CheckPasswordHash(user.Password, userBSON.Password) != nil {
		return "", errors.New("Wrong password")
	}

	// JWT payload
	data := map[string]interface{}{
		"_id": userBSON.ID,
	}

	// Generate user token
	token, err := util.GenerateToken(data)
	if err != nil {
		return "", errors.New("GenerateUserToken failed")
	}

	return token, nil

}
