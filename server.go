package main

import (
	"echo-app/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/users", service.CreateUser)
	e.GET("/users", service.AllUsers)
	e.GET("/users/:id", service.GetUserById)
	e.PUT("/users/:id", service.UpdateUserById)
	e.DELETE("/users/:id", service.DeleteUserById)

	e.Logger.Fatal(e.Start(":1323"))

}
