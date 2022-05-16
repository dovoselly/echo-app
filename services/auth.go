package services

import (
	"context"
	"echo-app/dao"
	"echo-app/database"
	"echo-app/models"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
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
		return models.UserBSON{}, errors.New(" da ton tai")
	}

	if errExist == nil {
		return models.UserBSON{}, errExist
	}

	// HashPassword
	payload.Password, _ = HashPassword(payload.Password)

	//Create user
	doc, err := dao.UserRegister(payload.ConvertToBSON())
	if err != nil {
		err = errors.New("khong the tao user")
		return doc, err
	}

	return doc, err

}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
