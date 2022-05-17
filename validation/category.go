package validation

import (
	"echo-app/models"
	"echo-app/utils"

	"github.com/labstack/echo/v4"
)

func CategoryCreateBody(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body models.CategoryCreateBody

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

		c.Set("body", body)

		return next(c)
	}
}

func CategoryUpdateBody(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body models.CategoryUpdateBody

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

		c.Set("body", body)

		return next(c)
	}
}
