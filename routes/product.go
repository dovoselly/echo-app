package routes

import (
	"echo-app/controllers"
	"echo-app/validations"
	"github.com/labstack/echo/v4"
)

func product(e *echo.Echo) {
	productRoutes := e.Group("/products")
	{
		productRoutes.GET("", controllers.ListProduct, validations.ListProduct)
		productRoutes.GET("/:id", controllers.ProductDetail)
		productRoutes.GET("/:id/reviews", controllers.ProductDetail)
	}
}
