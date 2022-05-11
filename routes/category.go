package routes

import (
	"echo-app/config"
	"echo-app/controllers"
	"echo-app/validations"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var envVars = config.GetEnv()

func category(e *echo.Echo) {
	isLogin := middleware.JWT([]byte(envVars.Jwt.SecretKey))
	categoryRouter := e.Group("/admin/categories", isLogin)
	{
		categoryRouter.POST("", controllers.CreateCategory, validations.CategoryCreateBody)
		categoryRouter.GET("", controllers.GetListCategory)
		categoryRouter.GET("/:id", controllers.GetCategoryByID)
		categoryRouter.PUT("/:id", controllers.UpdateCategory)
		categoryRouter.PATCH("/:id/disable", controllers.DisabledCategory)
		categoryRouter.PATCH("/:id/enabled", controllers.EnabledCategory)

	}
}
