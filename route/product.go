package route

import (
	"echo-app/controller"
	"echo-app/validation"

	"github.com/labstack/echo/v4"
)

func product(e *echo.Echo) {
	p := e.Group("/products")
	p.POST("", controller.CreateProduct)
	p.GET("", controller.ListProduct, validation.ListProduct)
	p.GET("/:id", controller.ProductDetail)
}
