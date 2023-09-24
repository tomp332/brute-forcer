package handlers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	manager "github.com/tomp332/gobrute/src/ manager"
	"net/http"
)

// HealthCheck
// @Summary API Health Check
// @Description API HealthCheck
// @Tags Health Check
// @Produce json
// @Success 200 {object} managerTypes.Health
// @Router /health [get]
func HealthCheck(c echo.Context) error {
	health, err := json.Marshal(&manager.ServiceHealth)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSONBlob(http.StatusOK, health)
}
