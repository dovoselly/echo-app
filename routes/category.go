package routes

import (
	"echo-app/config"
	"github.com/labstack/echo/v4"
)

var envVars = config.GetEnv()

func category(e *echo.Echo) {
	//isLogin := middleware.JWT([]byte(envVars.Jwt.SecretKey))
	//categoryRouter := e.Group("/admin/categories", isLogin)
	{
		//categoryRouter.POST("", controllers.CreateCategory)
		//categoryRouter.GET("", controllers.GetListCategory)
		//categoryRouter.GET("/:id", controllers.GetCategoryByID)
		//categoryRouter.PUT("/:id", controllers.UpdateCategory)
		//categoryRouter.PATCH("/:id/disable", controllers.D )

	}
}
