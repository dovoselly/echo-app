package route

import (
	"echo-app/config"
	"echo-app/controller"
	"echo-app/validation"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func category(e *echo.Echo) {
	var env = config.GetEnv()

	c := e.Group("/admin/categories")

	c.Use(middleware.JWT([]byte(env.Jwt.SecretKey)))

	c.POST("", controller.CreateCategory, validation.CategoryCreateBody)
	c.GET("", controller.GetListCategory)
	c.GET("/:id", controller.GetCategoryByID, validation.ValidateID)

	c.PUT("/:id", controller.UpdateCategoryByID, validation.ValidateID, validation.CategoryUpdateBody)
	c.DELETE("/:id", controller.DeleteCategoryByID, validation.ValidateID)
	c.PATCH("/:id/disable", controller.DisabledCategory)
	c.PATCH("/:id/enabled", controller.EnabledCategory)
}
