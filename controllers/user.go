package controllers

import (
	"echo-app/middlewares"
	"echo-app/models"
	"echo-app/services"
	"echo-app/utils"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
	"sync"
)

func Register(c echo.Context) error {
	// get body
	body := c.Get("body")
	payload, ok := body.(models.UserRegister)
	if !ok {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: utils.InvalidData,
		})
	}
	// check exist user
	foundUser := services.GetUserByEmail(payload.Email, payload.Username)
	if payload.Email == foundUser.Email {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: utils.EmailOrUsernameIsAlreadyExist,
		})
	}

	// init insert data
	id := primitive.NewObjectID()
	HashedPassword, err := hashPassword(payload.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: utils.InvalidData,
		})
	}
	insertData := models.User{
		ID:             id,
		Email:          strings.ToLower(payload.Email),
		Username:       strings.ToLower(payload.Username),
		HashedPassword: HashedPassword,
		Name:           payload.Name,
		DateOfBirth:    payload.DateOfBirth,
		Gender:         payload.Gender,
		CreatedAt:      utils.CurrentDateTime(),
		UpdatedAt:      utils.CurrentDateTime(),
	}

	// create user
	var wg sync.WaitGroup
	wg.Add(2)
	result := services.Register(insertData)
	return c.JSON(http.StatusOK, utils.Response{
		Message: result,
	})
}

func Login(c echo.Context) error {
	body := c.Get("body")
	payload, ok := body.(models.UserLogin)
	if !ok {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: utils.InvalidData,
		})
	}

	//check password
	foundUser := services.GetUserByEmail(payload.Email)
	fmt.Println(strings.ToLower(payload.Email), foundUser.Email, checkPasswordHash(payload.Password, foundUser.Password), payload.Password)

	if strings.ToLower(payload.Email) != foundUser.Email || !checkPasswordHash(payload.Password, foundUser.Password) {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: utils.EmailOrPasswordWrong,
		})
	}

	// Generate token
	fmt.Println(foundUser.ID)
	token, err := middlewares.GenerateToken(foundUser.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: utils.InvalidData,
		})
	}
	return c.JSON(http.StatusBadRequest, utils.Response{
		Message: utils.LoginSuccessFully,
		Token:   token,
	})
}

func ChangePassword(c echo.Context) error {

	return c.JSON(http.StatusOK, "")
}

func ResetPassword(c echo.Context) error {

	return c.JSON(http.StatusOK, "")
}

func GetUserByUsername(c echo.Context) error {

	return c.JSON(http.StatusOK, "")
}

func UpdateUserInfo(c echo.Context) error {

	return c.JSON(http.StatusOK, "")
}

func ChangeAvatar(c echo.Context) error {

	return c.JSON(http.StatusOK, "")
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
