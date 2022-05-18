package validation

import (
	"echo-app/models"
	"echo-app/util"

	"github.com/labstack/echo/v4"
)

func CreateReply(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body models.CreateReply
		if err := c.Bind(&body); err != nil {
			return util.Response400(c, nil, util.InvalidData)
		}

		if err := body.Validate(); err != nil {
			return util.Response400(c, nil, util.InvalidData)
		}

		c.Set("body", body)
		return next(c)
	}
}
