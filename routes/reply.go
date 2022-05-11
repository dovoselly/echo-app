package routes

import (
	"echo-app/controllers"
	"echo-app/validations"
	"github.com/labstack/echo/v4"
)

func reply(e *echo.Echo) {
	reviewRoutes := e.Group("/reviews")
	replyRoutes := e.Group("/replies")
	{
		reviewRoutes.POST("/:id", controllers.CreateReply)
	}
	{
		replyRoutes.GET("/:id", controllers.ListProduct, validations.ListProduct)
	}
}
