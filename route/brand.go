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

	b.POST("", brandCtrl.Create, brandVal.BrandCreateBody)
	b.GET("", brandCtrl.GetList)
	b.GET("/:id", brandCtrl.GetByID, validation.ValidateID)

	b.PUT("/:id", brandCtrl.UpdateByID, validation.ValidateID, brandVal.BrandUpdateBody)
	b.DELETE("/:id", brandCtrl.DeleteByID, validation.ValidateID)
	b.PATCH("/:id/disable", brandCtrl.Disabled)
	b.PATCH("/:id/enabled", brandCtrl.Enabled)
}
