package controller

// import (
// 	"echo-app/middleware"
// 	"echo-app/model"
// 	"echo-app/service"
// 	"echo-app/util"
// 	"fmt"
// 	"github.com/labstack/echo/v4"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"golang.org/x/crypto/bcrypt"
// 	"net/http"
// 	"strings"
// )

// type Query struct {
// 	Page  int `query:"page"`
// 	Limit int `query:"limit"`
// }

// func CreateUser(c echo.Context) error {
// 	// get body
// 	body := c.Get("body")
// 	payload, ok := body.(model.User)
// 	if !ok {
// 		return c.JSON(http.StatusBadRequest, util.Response{
// 			Message: util.InvalidData,
// 		})
// 	}
// 	// check exist user
// 	foundUser := service.GetUserByEmail(payload.Email)
// 	fmt.Println(payload, foundUser)
// 	if payload.Email == foundUser.Email {
// 		return c.JSON(http.StatusBadRequest, util.Response{
// 			Message: "Email is already in exist",
// 		})
// 	}

// 	// init insert data
// 	password, err := hashPassword(payload.Password)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, util.Response{
// 			Message: util.InvalidData,
// 		})
// 	}
// 	insertData := model.User{
// 		ID:        primitive.NewObjectID(),
// 		Name:      payload.Name,
// 		Dob:       payload.Dob,
// 		Email:     strings.ToLower(payload.Email),
// 		Password:  password,
// 		CreatedAt: util.CurrentDateTime(),
// 	}

// 	// create user
// 	result := service.CreateUser(insertData)
// 	return c.JSON(http.StatusOK, util.Response{
// 		Message: result,
// 	})
// }

// func Login(c echo.Context) error {
// 	body := c.Get("body")
// 	payload, ok := body.(model.UserLogin)
// 	if !ok {
// 		return c.JSON(http.StatusBadRequest, util.Response{
// 			Message: util.InvalidData,
// 		})
// 	}

// 	//check password
// 	foundUser := service.GetUserByEmail(payload.Email)
// 	fmt.Println(strings.ToLower(payload.Email), foundUser.Email, checkPasswordHash(payload.Password, foundUser.Password), payload.Password)

// 	if strings.ToLower(payload.Email) != foundUser.Email || !checkPasswordHash(payload.Password, foundUser.Password) {
// 		return c.JSON(http.StatusBadRequest, util.Response{
// 			Message: util.EmailOrPasswordWrong,
// 		})
// 	}

// 	// Generate token
// 	fmt.Println(foundUser.ID)
// 	token, err := middleware.GenerateToken(foundUser.ID)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, util.Response{
// 			Message: util.InvalidData,
// 		})
// 	}
// 	return c.JSON(http.StatusBadRequest, util.Response{
// 		Message: util.LoginSuccessFully,
// 		Token:   token,
// 	})
// }

// func AllUsers(c echo.Context) error {
// 	// get query param
// 	var query Query
// 	if err := c.Bind(&query); err != nil {
// 		return c.JSON(http.StatusOK, util.Response{
// 			Message: err.Error(),
// 		})
// 	}

// 	allUsers := service.AllUsers(query.Page, query.Limit)
// 	if allUsers == nil {
// 		allUsers = make([]model.User, 0)
// 	}
// 	return c.JSON(http.StatusOK, allUsers)
// }

// func GetUserById(c echo.Context) error {
// 	id, err := primitive.ObjectIDFromHex(c.Get("id").(string))
// 	if err != nil {
// 		return c.JSON(http.StatusOK, util.Response{
// 			Message: util.InvalidData,
// 		})
// 	}
// 	foundUser := service.GetUserById(id)
// 	return c.JSON(http.StatusOK, foundUser)
// }

// func UpdateUserById(c echo.Context) error {
// 	// bind data
// 	payload, ok := c.Get("body").(model.UserUpdate)
// 	if !ok {
// 		return c.JSON(http.StatusBadRequest, util.Response{
// 			Message: util.InvalidData,
// 		})
// 	}

// 	// parse id from path param
// 	id := getId(c)

// 	// init insert data
// 	insertData := model.UserUpdate{
// 		Name: payload.Name,
// 		Dob:  payload.Dob,
// 	}
// 	result := service.UpdateUserById(id, insertData)
// 	return c.JSON(http.StatusOK, util.Response{
// 		Message: result,
// 	})
// }

// func DeleteUserById(c echo.Context) error {
// 	id := getId(c)

// 	result := service.DeleteUserById(id)

// 	return c.JSON(http.StatusOK, util.Response{
// 		Message: result,
// 	})
// }

// func hashPassword(password string) (string, error) {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
// 	return string(bytes), err
// }

// func checkPasswordHash(password, hash string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// 	return err == nil
// }

// func getId(c echo.Context) primitive.ObjectID {
// 	authId := c.Get("id")
// 	authIdStrConv, ok := authId.(string)
// 	if !ok {
// 		if !ok {
// 			fmt.Println("Parse id to string failed")
// 		}
// 	}
// 	id, err := primitive.ObjectIDFromHex(authIdStrConv)
// 	if err != nil {
// 		fmt.Println("Parse string id to ObjectId failed")
// 	}
// 	return id
// }
