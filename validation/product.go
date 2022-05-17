package validation

import (
	"echo-app/models"
	"echo-app/utils"

	"github.com/labstack/echo/v4"
)

func ListProduct(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var query models.ProductQuery
		if err := c.Bind(&query); err != nil {
			return utils.Response400(c, nil, utils.InvalidData)
		}
		c.Set("query", query)
		return next(c)
	}
}
