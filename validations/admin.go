package validations

import (
	"echo-app/models"
	"echo-app/utils"
	"github.com/labstack/echo/v4"
)

// AdminLoginBody ...
func AdminLoginBody(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body models.AdminLoginBody

		// bind request data
		if err := c.Bind(&body); err != nil {
			if err != nil {
				return utils.Response400(c, nil, err.Error())
			}
		}

		// validate
		if err := body.Validate(); err != nil {
			return utils.Response400(c, nil, err.Error())
		}

		c.Set("adminLoginBody", body)

		return next(c)
	}
}

func ValidateAdminUpdateBody(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var admin models.Admin

		// bind request body
		err := c.Bind(&admin)

		if err != nil {
			return utils.Response400(c, nil, err.Error())
		}

		// success
		c.Set("adminRequestBody", admin)

		return next(c)
	}
}
