package validation

import (
	"echo-app/models"
	"echo-app/util"

	"github.com/labstack/echo/v4"
)

func CreateOrder(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			body models.OrderCreate
		)

		// Validate
		if err := c.Bind(&body); err != nil {
			return util.Response400(c, nil, err.Error())
		}

		// Success
		c.Set("body", body)
		return next(c)
	}
}
