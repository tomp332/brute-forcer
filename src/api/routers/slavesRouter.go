package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/tomp332/bruteForcer/src/api/handlers"
)

var SlaveRouter *echo.Group

func InitSlavesRouter() {
	SlaveRouter.POST("", handlers.AddSlaves)
	SlaveRouter.GET("", handlers.GetSlaves)
}
