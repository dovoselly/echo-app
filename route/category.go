package route

import (
	"echo-app/config"
	"echo-app/controller"
	"echo-app/validation"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var envVars = config.GetEnv()

func category(e *echo.Echo) {
	isLogin := middleware.JWT([]byte(envVars.Jwt.SecretKey))
	categoryRouter := e.Group("/admin/categories", isLogin)
	{
		categoryRouter.POST("", controller.CreateCategory, validation.CategoryCreateBody)
		categoryRouter.GET("", controller.GetListCategory)
		categoryRouter.GET("/:id", controller.GetCategoryByID, validation.ValidateID)

		categoryRouter.PUT("/:id", controller.UpdateCategoryByID, validation.ValidateID, validation.CategoryUpdateBody)
		categoryRouter.DELETE("/:id", controller.DeleteCategoryByID, validation.ValidateID)
		categoryRouter.PATCH("/:id/disable", controller.DisabledCategory)
		categoryRouter.PATCH("/:id/enabled", controller.EnabledCategory)

	}
}
