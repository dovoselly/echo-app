package route

import (
	"echo-app/config"
	"echo-app/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func user(e *echo.Echo) {
	var (
		env  = config.GetEnv()
		user = e.Group("/user")
	)

	user.POST("/register", controller.Register, userVal.Register)
	user.POST("/login", controller.Login, userVal.Login)

	user.Use(middleware.JWT([]byte(env.Jwt.SecretKey)))

	user.GET("/me", userCtrl.GetInfo)
	user.PUT("/me", userCtrl.UpdateInfo, userVal.UpdateInfo)
	user.PATCH("/me/password", userCtrl.ChangePassword, userVal.ChangePassword)

}
