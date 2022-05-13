package routes

import (
	"echo-app/config"
	"echo-app/controllers"
	"echo-app/validations"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func user(e *echo.Echo) {
	var env = config.GetEnv()

	user := e.Group("/user")

	user.POST("/register", controllers.Register, validations.UserRegister)
	user.POST("/login", controllers.Login, validations.UserLogin)

	user.Use(middleware.JWT([]byte(env.Jwt.SecretKey)))

	user.GET("/me", controllers.GetUserInfo)
	user.PUT("/me", controllers.UpdateUserInfo, validations.UpdateUerInfo)
	user.PATCH("/me/password", controllers.ChangeUserPassword, validations.ChangeUserPassword)

}
