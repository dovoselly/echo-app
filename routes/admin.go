package routes

import (
	"echo-app/controller"
	"echo-app/middleware"
	"github.com/labstack/echo/v4"
)

func admin(e echo.Echo) {
	adminRouter := e.Group("/admin")
	{
		adminRouter.POST("/register", controller.CreateUser, middleware.CreateUser)
		adminRouter.POST("/login", controller.Login, middleware.Login)
	}
}
