package route

import (
	"echo-app/config"
	"echo-app/controller"
	"echo-app/validation"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//var envVars = config.GetEnv()

func brand(e *echo.Echo) {
	var env = config.GetEnv()

	b := e.Group("/admin/brands")

	b.Use(middleware.JWT([]byte(env.Jwt.SecretKey)))

	b.POST("", controller.CreateBrand, validation.BrandCreateBody)
	b.GET("", controller.GetListBrand)
	b.GET("/:id", controller.GetBrandByID, validation.ValidateID)

	b.PUT("/:id", controller.UpdateCategoryByID, validation.ValidateID, validation.BrandUpdateBody)
	b.DELETE("/:id", controller.DeleteBrandByID, validation.ValidateID)
	b.PATCH("/:id/disable", controller.DisabledBrand)
	b.PATCH("/:id/enabled", controller.EnabledBrand)
}
