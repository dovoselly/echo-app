package validation

import (
	"echo-app/model"
	"echo-app/utils"

	"github.com/labstack/echo/v4"
)

func CreateOrder(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			body model.OrderCreate
		)

		// Validate
		if err := c.Bind(&body); err != nil {
			return utils.Response400(c, nil, err.Error())
		}

		// Success
		c.Set("body", body)
		return next(c)
	}
}
