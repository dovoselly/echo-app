package controller

import (
	"echo-app/middleware"
	"echo-app/model"
	"echo-app/service"
	"echo-app/util"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

func CreateUser(c echo.Context) error {
	// bind data
	payload, err := bindAndValidate(c)
	if err != nil {
		return err
	}

	// check exist user
	foundUser, err := service.GetUserByEmail(c, payload.Email)
	if err != nil {
		return err
	}
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
		Name:     payload.Name,
		Age:      payload.Age,
		Email:    payload.Email,
		Password: password,
	}

	// role admin
	path := c.Request().URL.Path
	if path == "/users/register-admin" {
		return service.CreateUser(c, &insertData, true)
	}

	// create user
	return service.CreateUser(c, &insertData, false)
}

func Login(c echo.Context) error {
	payload, err := bindAndValidate(c)
	if err != nil {
		return err
	}

	//check user
	foundUser, err := service.GetUserByEmail(c, payload.Email)
	if err != nil {
		return err
	}
	if payload.Email != foundUser.Email || !checkPasswordHash(payload.Password, foundUser.Password) {
		return c.JSON(http.StatusBadRequest, util.Response{
			Message: "Email or Password was wrong",
		})
	}

	// Generate token
	token, err := middleware.GenerateToken(foundUser.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusBadRequest, util.Response{
		Message: "Login successfully",
		Token:   token,
	})
}

func AllUsers(c echo.Context) error {
	// get query param
	page, err := strconv.ParseInt(c.QueryParam("page"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	limit, err := strconv.ParseInt(c.QueryParam("limit"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return service.AllUsers(c, page, limit)
}

func GetUserById(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return service.GetUserById(c, id)
}

func UpdateUserById(c echo.Context) error {
	// bind data
	payload, err := bindAndValidate(c)
	if err != nil {
		return err
	}

	// parse id from path param
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// init insert data
	insertData := model.User{
		Name: payload.Name,
		Age:  payload.Age,
	}

	return service.UpdateUserById(c, id, insertData)
}

func DeleteUserById(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
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

func bindAndValidate(c echo.Context) (*model.User, error) {
	payload := new(model.User)
	if err := c.Bind(payload); err != nil {
		return nil, c.JSON(http.StatusBadRequest, util.Response{
			Message: "invalid value",
		})
	}

	if !middleware.Validate(payload) {
		return nil, c.JSON(http.StatusBadRequest, util.Response{
			Message: "invalid value",
		})
	}
	return payload, nil
}
