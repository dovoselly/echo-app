package route

import (
	"echo-app/controller"
	"echo-app/validation"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//var envVars = config.GetEnv()

func brand(e *echo.Echo) {
	isLogin := middleware.JWT([]byte(envVars.Jwt.SecretKey))
	brandRouter := e.Group("/admin/brands", isLogin)
	{
		brandRouter.POST("", controller.CreateBrand, validation.BrandCreateBody)
		brandRouter.GET("", controller.GetListBrand)
		brandRouter.GET("/:id", controller.GetBrandByID, validation.ValidateID)
		brandRouter.PUT("/:id", controller.UpdateBrand, validation.ValidateID, validation.BrandUpdateBody)
		brandRouter.DELETE("/:id", controller.DeleteBrandByID, validation.ValidateID)
		brandRouter.PATCH("/:id/disable", controller.DisabledBrand)
		brandRouter.PATCH("/:id/enabled", controller.EnabledBrand)
	}
}
