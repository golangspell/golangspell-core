package controller

import (
	"net/http"

	"{{.ModuleName}}/domain"
	"github.com/labstack/echo"
)

//CheckHealth handles the application Health Check
func CheckHealth(c echo.Context) error {
	health := domain.Health{}
	health.Status = "UP"
	return c.JSON(http.StatusOK, health)
}
