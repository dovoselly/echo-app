package routes

import (
	"echo-app/controllers"
	"echo-app/validations"
	"github.com/labstack/echo/v4"
)

func review(e *echo.Echo) {
	productRoutes := e.Group("/products")
	reviewRoutes := e.Group("/reviews")
	{
		productRoutes.POST("/products/:id/reviews", controllers.CreateReview, validations.CreateReview)
		productRoutes.POST("/products/:id", controllers.ListReview, validations.ListReview)
	}
	{
		reviewRoutes.GET("", controllers.ListProduct, validations.ListProduct)
		reviewRoutes.GET("/:id", controllers.ProductDetail)
	}
}
