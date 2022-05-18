package service

import (
	"echo-app/model"
	"echo-app/utils"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct{}

func (u Auth) Register(body model.UserRegister) (primitive.ObjectID, error) {
	// HashPassword
	body.Password, _ = u.hashPassword(body.Password)

	//Create user
	id, err := userDAO.Register(body.ConvertToBSON())
	if err != nil {
		err = errors.New("khong the tao user")
		return id, err
	}

	return id, err
}

func (u Auth) Login(user model.UserLogin) (string, error) {
	// get user
	userBSON, err := userDAO.GetByUsername(user.Username)
	if err != nil {
		return "", errors.New(utils.NOT_EXIST_USER)
	}

	// verify user password
	if u.checkPasswordHash(user.Password, userBSON.Password) != nil {
		return "", errors.New(utils.WRONG_PASSWORD)
	}

	// JWT payload
	data := map[string]interface{}{
		"_id": userBSON.Id,
	}

	// Generate user token
	token, err := utils.GenerateToken(data)
	if err != nil {
		return "", errors.New(utils.GENERATE_TOKEN_FAILED)
	}

	return token, nil
}

func (u Auth) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (u Auth) checkPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}
