package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/tomp332/gobrute/pkg/manager/api/handlers"
)

var BruteRouter *echo.Group

type BruteRoute struct{}

func (r BruteRoute) InitRouter() error {
	BruteRouter.POST("", handlers.StartBruteForce)
	return nil
}
