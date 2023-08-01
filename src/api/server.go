package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
	"github.com/tomp332/bruteForcer/src"
	"github.com/tomp332/bruteForcer/src/api/middlewares"
	"github.com/tomp332/bruteForcer/src/api/routers"
)

var MainServerObj *echo.Echo

func InitServer() {
	MainServerObj = echo.New()
	// Middlewares
	MainServerObj.Use(middlewares.LogRequest)

	// Groups
	MainServerObj.GET("/swagger/*", echoSwagger.WrapHandler)
	v1ApiGroup := MainServerObj.Group("/api/v1")
	routers.MainRouter = v1ApiGroup.Group("")
	routers.SlaveRouter = v1ApiGroup.Group("/slaves")
	routers.CredsRouter = v1ApiGroup.Group("/creds")

	// Routers
	err := routers.MainRoute{}.InitRouter()
	if err != nil {
		return
	}
	err = routers.CredsRoute{}.InitRouter()
	if err != nil {
		return
	}
	err = routers.SlaveRoute{}.InitRouter()
	if err != nil {
		return
	}

	// Start server
	addr := fmt.Sprintf("%s:%d", src.GlobalSettings.ServerHost, src.GlobalSettings.ServerPort)
	err = MainServerObj.Start(addr)
	if err != nil {
		panic(err)
	}
}
