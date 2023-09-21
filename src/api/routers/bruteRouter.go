package routers

import (
	"github.com/labstack/echo/v4"
)

var BruteRouter *echo.Group

type BruteRoute struct{}

func (r BruteRoute) InitRouter() error {
	return nil
}
