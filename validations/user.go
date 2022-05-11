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
			payload models.UserRegister
		)

		// Validate
		c.Bind(&payload)
		err := payload.Validate()

		//if err
		if err != nil {
			return utils.Response400(c, nil, err.Error())
		}

		// Success
		c.Set("payload", payload)
		return next(c)
	}
}
