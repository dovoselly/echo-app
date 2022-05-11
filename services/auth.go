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
		user       *models.UserResonse
	)

	// // Check exist username, password
	// errExistUsername := collection.FindOne(ctx, bson.M{"username": payload.Username}).Decode(&user)
	// errExistEmail := collection.FindOne(ctx, bson.M{"email": payload.Email}).Decode(&user)

	// if errExistUsername == nil && errExistEmail == nil {
	// 	err := errors.New("username and email da ton tai")
	// 	return models.UserBSON{}, err
	// } else if errExistUsername == nil {
	// 	err := errors.New("Username da ton tai")
	// 	return models.UserBSON{}, err
	// } else if errExistEmail == nil {
	// 	err := errors.New("Email da ton tai")
	// 	return models.UserBSON{}, err
	// }

	fmt.Println("user11111: ", user)
	errExist := collection.FindOne(ctx, bson.M{"$or": []bson.M{{"email": payload.Email}, {"username": payload.Username}}}).Decode(&user)
	fmt.Println("user: ", user)
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
