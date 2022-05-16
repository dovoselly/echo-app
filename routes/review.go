package routes

import (
	"echo-app/config"
	"echo-app/controllers"
	"echo-app/validations"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func review(e *echo.Echo) {
	productRoutes := e.Group("/products", middleware.JWT([]byte(config.GetEnv().Jwt.SecretKey)))
	reviewRoutes := e.Group("/reviews")
	{
		productRoutes.POST("/:id/reviews", controllers.CreateReview, validations.CreateReview)
		productRoutes.POST("/:id", controllers.ListReview, validations.ListReview)
	}
	{
		reviewRoutes.GET("", controllers.ListReview, validations.ListReview)
	}
}
