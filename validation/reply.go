package validation

import (
	"echo-app/models"
	"echo-app/utils"

	"github.com/labstack/echo/v4"
)

func CreateReply(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body models.CreateReply
		if err := c.Bind(&body); err != nil {
			return utils.Response400(c, nil, utils.InvalidData)
		}

		if err := body.Validate(); err != nil {
			return utils.Response400(c, nil, utils.InvalidData)
		}

		c.Set("body", body)
		return next(c)
	}
}
