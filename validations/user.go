package validations

import (
	"echo-app/models"
	"echo-app/utils"

	"github.com/labstack/echo/v4"
)

// PlayerCreate ...
func UserRegister(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			body models.UserRegister
		)

		// Validate
		c.Bind(&body)
		err := body.Validate()

		//if err
		if err != nil {
			return utils.Response400(c, nil, err.Error())
		}

		// Success
		c.Set("body", body)
		return next(c)
	}
}

func UserLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			body models.UserLogin
		)

		// Validate
		c.Bind(&body)
		err := body.Validate()

		if err != nil {
			return utils.Response400(c, nil, err.Error())
		}
		// Success
		c.Set("body", body)
		return next(c)
	}
}

func UserChangePassword(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// code
		var (
			body models.UserChangePassword
		)

		//validate
		c.Bind(&body)

		err := body.Validate()

		// if err
		if err != nil {
			return utils.Response400(c, nil, err.Error())
		}

		// Success
		c.Set("body", body)
		return next(c)
	}
}

func IDUserInToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		// GetJWTPaylaod
		jwtPayload, err := utils.GetJWTPayload(c)

		if err != nil {
			return utils.Response400(c, nil, err.Error())
		}

		id := jwtPayload["id"].(string)

		// ValidateObjectID
		if err := utils.ValidateObjectID(id); err != nil {
			return utils.Response400(c, nil, err.Error())
		}

		c.Set("id", id)
		return next(c)
	}
}
