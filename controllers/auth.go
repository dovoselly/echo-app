package controllers

import (
	"echo-app/models"
	"echo-app/services"
	"echo-app/utils"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func Register(c echo.Context) error {

	var (
		payload = c.Get("payload").(models.UserRegister)
	)

	// Process data
	rawData, err := services.UserRegister(payload)

	if err != nil {

		return utils.Response400(c, nil, err.Error())
	}

	// Success
	return utils.Response200(c, bson.M{
		"_id":       rawData.ID,
		"createdAt": rawData.CreatedAt,
	}, "")
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
