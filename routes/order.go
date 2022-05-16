package routes

import (
	"echo-app/config"
	"echo-app/controllers"
	"echo-app/validations"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func order(e *echo.Echo) {
	var env = config.GetEnv()

	user := e.Group("/orders")

	user.Use(middleware.JWT([]byte(env.Jwt.SecretKey)))

	user.GET("", controllers.GetAllOrdersByUserId)
	user.POST("", controllers.CreateOrder, validations.CreateOrder)

}
