package route

import (
	"echo-app/config"
	"echo-app/controller"
	"echo-app/validation"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func reply(e *echo.Echo) {
	reviewRoutes := e.Group("/reviews", middleware.JWT([]byte(config.GetEnv().Jwt.SecretKey)))
	{
		reviewRoutes.POST("/:id", controller.CreateReply, validation.CreateReply)
	}

	replyRoutes := e.Group("/replies", middleware.JWT([]byte(config.GetEnv().Jwt.SecretKey)))
	{
		replyRoutes.PUT("/:id", controller.UpdateReply, validation.CreateReply)
		replyRoutes.DELETE("/:id", controller.DeleteReply)
	}
}
