package main

import (
	"echo-app/controller"
	"echo-app/database"
	midd "echo-app/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	database.RunDB()

	usersRouter := e.Group("/users")
	{
		usersRouter.POST("/register", controller.CreateUser)
		usersRouter.POST("/register-admin", controller.CreateUser)
		usersRouter.POST("/login", controller.Login)
		usersRouter.GET("", controller.AllUsers)
		usersRouter.GET("/:id", controller.GetUserById)
		usersRouter.PUT("/:id", controller.UpdateUserById, midd.Auth)
		usersRouter.DELETE("/:id", controller.DeleteUserById, midd.Auth)
	}

	e.Logger.Fatal(e.Start(":1323"))

}
