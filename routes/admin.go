package routes

import (
	"echo-app/controllers"
	"echo-app/validations"
	"github.com/labstack/echo/v4"
)

func admin(e echo.Echo) {
	adminRouter := e.Group("/admin")
	{
		adminRouter.POST("/admin-login", controllers.AdminLogin, validations.AdminLoginBody)
		adminRouter.GET("/me", controllers.MyProfileAdmin)
		adminRouter.PUT("/me", controllers.UpdateMyProfileAdmin)
		adminRouter.PATCH("/me/password", controllers.ChangePasswordAdmin)
		adminRouter.PATCH("/me/avatar", controllers.ChangeAvatarAdmin)
	}
}
