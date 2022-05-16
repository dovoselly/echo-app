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
		var body models.CreateReview
		if err := c.Bind(&body); err != nil {
			return utils.Response400(c, nil, err.Error())
		}
		if err := body.Validate(); err != nil {
			return utils.Response400(c, nil, err.Error())
		}
		if body.Rating > 5 || body.Rating < 1 {
			return utils.Response400(c, nil, "rating must be between 1 - 5")
		}
		c.Set("body", body)
		return next(c)
	}
}
