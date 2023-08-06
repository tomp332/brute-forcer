package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/tomp332/gospray/src/api/handlers"
)

var SlaveRouter *echo.Group

type SlaveRoute struct{}

func (r SlaveRoute) InitRouter() error {
	SlaveRouter.POST("", handlers.AddSlaves)
	SlaveRouter.GET("", handlers.GetSlaves)
	return nil
}
