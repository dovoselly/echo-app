package route

import (
	"echo-app/config"
	"echo-app/controller"
	"echo-app/validation"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func order(e *echo.Echo) {
	var env = config.GetEnv()

	user := e.Group("/orders")

	user.Use(middleware.JWT([]byte(env.Jwt.SecretKey)))

	user.GET("", controller.GetAllOrdersByUserId)
	user.POST("", controller.CreateOrder, validation.CreateOrder)

}
