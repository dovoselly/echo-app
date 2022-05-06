package routes

import (
	"echo-app/controller"
	"echo-app/middleware"
	"github.com/labstack/echo/v4"
)

func user(e *echo.Echo) {
	e.POST("/register", controller.Register)
	e.POST("/login", controller.Login)
	e.PATCH("/password", controller.ChangePassword, middleware.Auth)
	e.POST("/reset-password", controller.ResetPassword, middleware.Auth)
	e.GET("/:username", controller.GetUserByUsername)
	e.PUT("/me", controller.UpdateUserInfo, middleware.Auth)
	e.PATCH("/me/change-avatar", controller.ChangeAvatar, middleware.Auth)
}
