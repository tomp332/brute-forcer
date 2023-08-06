package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/tomp332/gobrute/src/api/handlers"
)

var CredsRouter *echo.Group

type CredsRoute struct{}

func (r CredsRoute) InitRouter() error {
	CredsRouter.POST("", handlers.AddCreds)
	CredsRouter.GET("", handlers.GetCreds)
	CredsRouter.DELETE("", handlers.DeleteCreds)
	CredsRouter.PUT("", handlers.UpdateCreds)
	return nil
}
