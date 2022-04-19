package middleware

import (
	"echo-app/model"
	"echo-app/util"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user model.User
		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		if err := user.Validate(); err != nil {
			fmt.Println(err.Error(), "222222")
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		c.Set("body", user)
		return next(c)
	}
}

func Login(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user model.UserLogin
		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		if err := user.Validate(); err != nil {
			fmt.Println(err.Error(), "222222")
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		c.Set("body", user)
		return next(c)
	}
}

func UpdateUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var UserUpdate model.UserUpdate
		if err := c.Bind(&UserUpdate); err != nil {
			return c.JSON(http.StatusBadRequest, util.Response{
				Message: err.Error(),
			})
		}

		if err := UserUpdate.Validate(); err != nil {
			return c.JSON(http.StatusBadRequest, util.Response{
				Message: err.Error(),
			})
		}

		c.Set("body", UserUpdate)
		return next(c)
	}
}
