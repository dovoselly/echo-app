package route

import (
	"echo-app/controller"
	"echo-app/validation"

	"github.com/labstack/echo/v4"
)

func product(e *echo.Echo) {
	productRoutes := e.Group("/products")
	{
		productRoutes.GET("", controller.ListProduct, validation.ListProduct)
		productRoutes.GET("/:id", controller.ProductDetail)
	}
}
