package route

import (
	"github.com/labstack/echo/v4"
)

func product(e *echo.Echo) {
	var p = e.Group("/products")

	p.GET("", productCtrl.GetListProduct, productVal.GetListProduct)
	p.GET("/:id", productCtrl.GetProductDetail)
}
