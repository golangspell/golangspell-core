package controller

import (
	"net/http"

	"{{.ModuleName}}/config"
	"{{.ModuleName}}/domain"
	"github.com/labstack/echo"
)

//GetInfo of the application like version
func GetInfo(c echo.Context) error {
	info := domain.Info{Version: config.Version}
	return c.JSON(http.StatusOK, info)
}
