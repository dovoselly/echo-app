package validations

import (
	"echo-app/models"
	"echo-app/utils"
	"github.com/labstack/echo/v4"
)

func ListReview(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var query models.ReviewQuery
		if err := c.Bind(&query); err != nil {
			return utils.Response400(c, nil, utils.InvalidData)
		}
		c.Set("query", query)
		return next(c)
	}
}

func CreateReview(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload models.CreateReview
		if err := c.Bind(&payload); err != nil {
			return utils.Response400(c, nil, utils.InvalidData)
		}
		c.Set("payload", payload)
		return next(c)
	}
}
