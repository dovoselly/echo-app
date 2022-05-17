package validation

import (
	"echo-app/model"
	"echo-app/util"

	"github.com/labstack/echo/v4"
)

func ListProduct(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var query model.ProductQuery
		if err := c.Bind(&query); err != nil {
			return util.Response400(c, nil, util.InvalidData)
		}
		c.Set("query", query)
		return next(c)
	}
}
