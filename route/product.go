package route

import (
	"echo-app/controller"
	"echo-app/validation"
	"github.com/labstack/echo/v4"
)

func product(e *echo.Echo) {
	var (
		p = e.Group("/products")
		c = controller.Product{}
		v = validation.Product{}
	)

	p.GET("", c.GetListProduct, v.GetListProduct)
	p.GET("/:id", c.GetProductDetail)
}
