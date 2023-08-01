package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/tomp332/bruteForcer/src/api/handlers"
)

var MainRouter *echo.Group

type MainRoute struct{}

func (r MainRoute) InitRouter() error {
	MainRouter.GET("/health", handlers.HealthCheck)
	return nil
}
