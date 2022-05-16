package routes

import (
	"echo-app/config"
	"echo-app/controllers"
	"echo-app/validations"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func review(e *echo.Echo) {
	productRoutes := e.Group("/products")
	{
		productRoutes.POST("/:id/reviews", controllers.CreateReview, validations.CreateReview, middleware.JWT([]byte(config.GetEnv().Jwt.SecretKey)))
		productRoutes.GET("/:id/reviews", controllers.ListReview, validations.ListReview)
	}
}
