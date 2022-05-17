package route

import (
	"echo-app/config"
	"echo-app/controller"
	"echo-app/validation"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func review(e *echo.Echo) {
	productRoutes := e.Group("/products")
	{
		productRoutes.POST("/:id/reviews", controller.CreateReview, validation.CreateReview, middleware.JWT([]byte(config.GetEnv().Jwt.SecretKey)))
		productRoutes.GET("/:id/reviews", controller.ListReview, validation.ListReview)
	}
}
