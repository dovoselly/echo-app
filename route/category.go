package route

import (
	"echo-app/config"
	"echo-app/validation"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func category(e *echo.Echo) {
	var env = config.GetEnv()
	c := e.Group("/admin/categories")
	c.Use(middleware.JWT([]byte(env.Jwt.SecretKey)))

	c.POST("", categoryCtrl.CreateCategory, categoryVal.CreateBody)
	c.GET("", categoryCtrl.GetListCategory)
	c.GET("/:id", categoryCtrl.GetCategoryByID, validation.ValidateID)
	c.PUT("/:id", categoryCtrl.UpdateCategoryByID, validation.ValidateID, categoryVal.UpdateBody)
	c.DELETE("/:id", categoryCtrl.DeleteCategoryByID, validation.ValidateID)
	c.PATCH("/:id/disable", categoryCtrl.DisabledCategory)
	c.PATCH("/:id/enabled", categoryCtrl.EnabledCategory)
}
