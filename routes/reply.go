package routes

import (
	"echo-app/config"
	"echo-app/controllers"
	"echo-app/validations"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func reply(e *echo.Echo) {
	reviewRoutes := e.Group("/reviews", middleware.JWT([]byte(config.GetEnv().Jwt.SecretKey)))
	{
		reviewRoutes.POST("/:id", controllers.CreateReply, validations.CreateReply)
	}

	replyRoutes := e.Group("/replies", middleware.JWT([]byte(config.GetEnv().Jwt.SecretKey)))
	{
		replyRoutes.PUT("/:id", controllers.UpdateReply, validations.CreateReply)
		replyRoutes.DELETE("/:id", controllers.UpdateReply)
	}
}
