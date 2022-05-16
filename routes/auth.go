package routes

import (
	"github.com/labstack/echo/v4"
)

func auth(e *echo.Echo) {
<<<<<<< HEAD
=======
	//adminRouter.POST("/admin-login", controllers.AdminLogin, validations.AdminLoginBody)

	e.POST("/register", controllers.Register, validations.UserRegister)
	e.POST("/admin/admin-login", controllers.AdminLogin, validations.AdminLoginBody)
	// e.POST("/login", controllers.Login)
	// e.PATCH("/password", controllers.ChangePassword, middlewares.Auth)
	// e.POST("/reset-password", controllers.ResetPassword, middlewares.Auth)
	// e.GET("/:username", controllers.GetUserByUsername)
	// e.PUT("/me", controllers.UpdateUserInfo, middlewares.Auth)
	// e.PATCH("/me/change-avatar", controllers.ChangeAvatar, middlewares.Auth)
>>>>>>> 044f84be6e7624a9a1fbd2fe4f3a374f3eaa60df
}
