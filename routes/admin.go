package routes

import (
	"echo-app/controllers"
	"echo-app/validations"
)


func admin(e *echo.Echo) {
	isLogin := middleware.JWT([]byte(envVars.Jwt.SecretKey))
	adminRouter := e.Group("/admin", isLogin)
	{
		//adminRouter.POST("/admin-login", controllers.AdminLogin, validations.AdminLoginBody)
		adminRouter.GET("/me", controllers.MyProfileAdmin, isLogin)
		adminRouter.PUT("/me", controllers.UpdateMyProfileAdmin, isLogin, validations.ValidateAdminUpdateBody)
		adminRouter.PATCH("/me/password", controllers.ChangePasswordAdmin)
		adminRouter.PATCH("/me/avatar", controllers.ChangeAvatarAdmin)
	}
}