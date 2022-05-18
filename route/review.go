package route

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func review(e *echo.Echo) {
	var p = e.Group("/products")

	p.POST("/:id/reviews", reviewCtrl.CreateReview, reviewVal.CreateReview, middleware.JWT(secretKeyBytes))
	p.GET("/:id/reviews", reviewCtrl.GetListReview, reviewVal.GetListReview)
}
