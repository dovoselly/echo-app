package main

import (
	"echo-app/config"
	"echo-app/dao"
	"echo-app/database"
	"echo-app/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	config.InitDotEnv()
	database.Connect()
}

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.Use(middleware.Recover())
	config.InitDotEnv()
	database.Connect()

	dao.InitAdminAccount()

	routes.Routes(e)

	e.Logger.Fatal(e.Start(":" + config.GetEnv().Port))

}
