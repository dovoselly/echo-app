package route

import (
	"echo-app/config"
	"echo-app/validation"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//var envVars = config.GetEnv()

func brand(e *echo.Echo) {
	var env = config.GetEnv()

	b := e.Group("/admin/brands")

	b.Use(middleware.JWT([]byte(env.Jwt.SecretKey)))

	b.POST("", brandCtrl.CreateBrand, brandVal.BrandCreateBody)
	b.GET("", brandCtrl.GetListBrand)
	b.GET("/:id", brandCtrl.GetBrandByID, validation.ValidateID)

	b.PUT("/:id", brandCtrl.UpdateBrandByID, validation.ValidateID, brandVal.BrandUpdateBody)
	b.DELETE("/:id", brandCtrl.DeleteBrandByID, validation.ValidateID)
	b.PATCH("/:id/disable", brandCtrl.DisabledBrand)
	b.PATCH("/:id/enabled", brandCtrl.EnabledBrand)
}
