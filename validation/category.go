package validation

import (
	"echo-app/model"
	"echo-app/util"

	"github.com/labstack/echo/v4"
)

func CategoryCreateBody(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body model.CategoryCreateBody

		// bind request data
		if err := c.Bind(&body); err != nil {
			if err != nil {
				return util.Response400(c, nil, err.Error())
			}
		}

		// validate
		if err := body.Validate(); err != nil {
			return util.Response400(c, nil, err.Error())
		}

		c.Set("body", body)

		return next(c)
	}
}

func CategoryUpdateBody(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body model.CategoryUpdateBody

		// bind request data
		if err := c.Bind(&body); err != nil {
			if err != nil {
				return util.Response400(c, nil, err.Error())
			}
		}

		// validate
		if err := body.Validate(); err != nil {
			return util.Response400(c, nil, err.Error())
		}

		c.Set("body", body)

		return next(c)
	}
}
