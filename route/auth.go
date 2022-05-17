package route

import (
	"echo-app/controller"
	"echo-app/validation"

	"github.com/labstack/echo/v4"
)

func auth(e *echo.Echo) {
	e.POST("/admin/admin-login", controller.AdminLogin, validation.AdminLoginBody)
}
