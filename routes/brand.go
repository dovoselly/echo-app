package routes

import (
	"echo-app/controllers"
	"echo-app/validations"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//var envVars = config.GetEnv()

func brand(e *echo.Echo) {
	isLogin := middleware.JWT([]byte(envVars.Jwt.SecretKey))
	brandRouter := e.Group("/admin/brands", isLogin)
	{
		brandRouter.POST("", controllers.CreateBrand, validations.BrandCreateBody)
		brandRouter.GET("", controllers.GetListBrand)
		brandRouter.GET("/:id", controllers.GetBrandByID, validations.ValidateID)
		brandRouter.PUT("/:id", controllers.UpdateBrand, validations.ValidateID, validations.BrandUpdateBody)
		brandRouter.DELETE("/:id", controllers.DeleteBrandByID, validations.ValidateID)
		brandRouter.PATCH("/:id/disable", controllers.DisabledBrand)
		brandRouter.PATCH("/:id/enabled", controllers.EnabledBrand)
	}
}