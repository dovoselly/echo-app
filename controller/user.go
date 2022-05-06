package controller

import (
	"echo-app/middleware"
	"echo-app/model"
	"echo-app/service"
	"echo-app/util"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
)

func Register(c echo.Context) error {
	// get body
	body := c.Get("body")
	payload, ok := body.(model.User)
	if !ok {
		return c.JSON(http.StatusBadRequest, util.Response{
			Message: util.InvalidData,
		})
	}
	// check exist user
	foundUser := service.GetUserByEmail(payload.Email)
	if payload.Email == foundUser.Email {
		return c.JSON(http.StatusBadRequest, util.Response{
			Message: "Email is already in exist",
		})
	}

	// init insert data
	password, err := hashPassword(payload.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.Response{
			Message: util.InvalidData,
		})
	}
	insertData := model.User{
		ID:        primitive.NewObjectID(),
		Name:      payload.Name,
		Dob:       payload.Dob,
		Email:     strings.ToLower(payload.Email),
		Password:  password,
		CreatedAt: util.CurrentDateTime(),
	}

	// create user
	result := service.CreateUser(insertData)
	return c.JSON(http.StatusOK, util.Response{
		Message: result,
	})
}

func Login(c echo.Context) error {
	body := c.Get("body")
	payload, ok := body.(model.UserLogin)
	if !ok {
		return c.JSON(http.StatusBadRequest, util.Response{
			Message: util.InvalidData,
		})
	}

	//check password
	foundUser := service.GetUserByEmail(payload.Email)
	fmt.Println(strings.ToLower(payload.Email), foundUser.Email, checkPasswordHash(payload.Password, foundUser.Password), payload.Password)

	if strings.ToLower(payload.Email) != foundUser.Email || !checkPasswordHash(payload.Password, foundUser.Password) {
		return c.JSON(http.StatusBadRequest, util.Response{
			Message: util.EmailOrPasswordWrong,
		})
	}

	// Generate token
	fmt.Println(foundUser.ID)
	token, err := middleware.GenerateToken(foundUser.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.Response{
			Message: util.InvalidData,
		})
	}
	return c.JSON(http.StatusBadRequest, util.Response{
		Message: util.LoginSuccessFully,
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
