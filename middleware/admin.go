package middleware

import (
	"echo-app/utils"

	"github.com/labstack/echo/v4"
)

func CheckAdminRole(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		jwtPayload, _ := utils.GetJWTPayload(c)
		//adminID := jwtPayload["id"].(string)

		if jwtPayload["isAdmin"] == true {
			return next(c)
		}

		return utils.Response400(c, nil, "authorization fail: not admin")
	}
}
