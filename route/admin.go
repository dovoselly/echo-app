package route

import (
	"echo-app/controller"
	"echo-app/validation"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func admin(e *echo.Echo) {
	isLogin := middleware.JWT([]byte(envVars.Jwt.SecretKey))
	adminRouter := e.Group("/admin", isLogin)
	{
		//adminRouter.POST("/admin-login", controller.AdminLogin, validation.AdminLoginBody)
		adminRouter.GET("/me", controller.MyProfileAdmin, isLogin)
		adminRouter.PUT("/me", controller.UpdateMyProfileAdmin, isLogin, validation.ValidateAdminUpdateBody)
		adminRouter.PATCH("/me/password", controller.ChangePasswordAdmin)
		adminRouter.PATCH("/me/avatar", controller.ChangeAvatarAdmin)
	}
}
