package routes

import (
	"echo-app/controller"
	"echo-app/middleware"
	"github.com/labstack/echo/v4"
)

func user(e *echo.Echo) {
	usersRouter := e.Group("/users")
	{
		usersRouter.POST("/register", controller.CreateUser, middleware.CreateUser)
		usersRouter.POST("/login", controller.Login, middleware.Login)
		usersRouter.GET("", controller.AllUsers)
		usersRouter.GET("/", controller.GetUserById)
		usersRouter.PUT("/", controller.UpdateUserById, middleware.Auth, middleware.UpdateUser)
		usersRouter.DELETE("/", controller.DeleteUserById, middleware.Auth)
	}
}
