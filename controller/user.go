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
	"time"
)

type Query struct {
	Page  int `query:"page"`
	Limit int `query:"limit"`
}

func CreateUser(c echo.Context) error {
	// get body
	body := c.Get("body")
	payload, ok := body.(*model.User)
	if !ok {
		return c.JSON(http.StatusBadRequest, util.Response{
			Message: "Invalid data",
		})
	}
	// check exist user
	foundUser := service.GetUserByEmail(payload.Email)
	fmt.Println(payload, foundUser)
	if payload.Email == foundUser.Email {
		return c.JSON(http.StatusBadRequest, util.Response{
			Message: "Email is already in exist",
		})
	}

	// init insert data
	password, err := hashPassword(payload.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.Response{
			Message: "invalid value",
		})
	}
	insertData := model.User{
		ID:        primitive.NewObjectID(),
		Name:      payload.Name,
		Dob:       payload.Dob,
		Email:     strings.ToLower(payload.Email),
		Password:  password,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	// create user
	result := service.CreateUser(insertData)
	return c.JSON(http.StatusOK, util.Response{
		Message: result,
	})
}

func Login(c echo.Context) error {
	payload, ok := c.Get("body").(model.UserLogin)
	if ok {
		return c.JSON(http.StatusBadRequest, util.Response{
			Message: "Invalid data",
		})
	}

	//check password
	foundUser := service.GetUserByEmail(payload.Email)
	fmt.Println(strings.ToLower(payload.Email), foundUser.Email, checkPasswordHash(payload.Password, foundUser.Password), payload.Password)

	if strings.ToLower(payload.Email) != foundUser.Email || !checkPasswordHash(payload.Password, foundUser.Password) {
		return c.JSON(http.StatusBadRequest, util.Response{
			Message: "Email or Password was wrong",
		})
	}

	// Generate token
	token, err := middleware.GenerateToken(foundUser.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.Response{
			Message: "Invalid data",
		})
	}
	return c.JSON(http.StatusBadRequest, util.Response{
		Message: "Login successfully",
		Token:   token,
	})
}

func AllUsers(c echo.Context) error {
	// get query param
	var query Query
	if err := c.Bind(&query); err != nil {
		return c.JSON(http.StatusOK, util.Response{
			Message: err.Error(),
		})
	}

	allUsers := service.AllUsers(c, query.Page, query.Limit)
	if allUsers == nil {
		allUsers = make([]model.User, 0)
	}
	return c.JSON(http.StatusOK, allUsers)
}

func GetUserById(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.Response{
			Message: "Invalid data",
		})
	}

	return service.GetUserById(c, id)
}

func UpdateUserById(c echo.Context) error {
	// bind data
	payload, ok := c.Get("body").(model.UserUpdate)
	if !ok {
		return c.JSON(http.StatusBadRequest, util.Response{
			Message: "Invalid data",
		})
	}

	// parse id from path param
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.Response{
			Message: err.Error(),
		})
	}

	// init insert data
	insertData := model.User{
		Name: payload.Name,
		Dob:  payload.Dob,
	}

	return service.UpdateUserById(c, id, insertData)
}

func DeleteUserById(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.Response{
			Message: "Invalid data",
		})
	}

	return service.DeleteUserById(c, id)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
