package validation

import (
	"echo-app/models"
	"echo-app/util"

	"github.com/labstack/echo/v4"
)

type User struct{}

func (u User) Register(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			body models.UserRegister
		)
		// Validate
		c.Bind(&body)

		//if err
		if err := body.Validate(); err != nil {
			return util.Response400(c, nil, err.Error())
		}

		// Success
		c.Set("body", body)
		return next(c)
	}
}

func (u User) Login(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body models.UserLogin

		// Validate
		c.Bind(&body)

		if err := body.Validate(); err != nil {
			return util.Response400(c, nil, err.Error())
		}
		// Success
		c.Set("body", body)
		return next(c)
	}
}

func (u User) ChangePassword(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body models.UserChangePassword

		//validate
		c.Bind(&body)

		// if err
		if err := body.Validate(); err != nil {
			return util.Response400(c, nil, err.Error())
		}

		// Success
		c.Set("body", body)
		return next(c)
	}
}

func (u User) UpdateInfo(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body models.UserUpdate

		// validate
		c.Bind(&body)

		if err := body.Validate(); err != nil {
			return util.Response400(c, nil, err.Error())
		}

		// Success
		c.Set("body", body)
		return next(c)
	}
}
