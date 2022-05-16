package routes

import "github.com/labstack/echo/v4"

func Routes(e *echo.Echo) {
	auth(e)
	user(e)
	order(e)

}
