package main

import (
	"{{.ModuleName}}/config"
	"{{.ModuleName}}/controller"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	controller.MapRoutes(e)

	e.Logger.Fatal(e.Start(":" + config.Values.Port))
}
