package routes

import "github.com/labstack/echo/v4"

func Routes(e *echo.Echo) {
	user(e)
	product(e)
	//admin(e)
}
