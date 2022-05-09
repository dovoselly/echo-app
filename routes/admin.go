package routes

import (
	"echo-app/controllers"
	"echo-app/middlewares"
	"github.com/labstack/echo/v4"
)

func admin(e echo.Echo) {
	adminRouter := e.Group("/admin")
	{
		adminRouter.POST("/register", controllers.CreateUser, middlewares.CreateUser)
		adminRouter.POST("/login", controllers.Login, middlewares.Login)
	}
}
