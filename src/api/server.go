package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/tomp332/bruteForcer/src"
	"github.com/tomp332/bruteForcer/src/api/handlers"
	"github.com/tomp332/bruteForcer/src/api/middlewares"
	"github.com/tomp332/bruteForcer/src/api/routers"
)

var MainServerObj *echo.Echo

func InitServer() {
	MainServerObj = echo.New()
	// Middlewares
	MainServerObj.Use(middlewares.LogRequest)
	//MainServerObj.Pre(middleware.RemoveTrailingSlash())

	// Groups
	routers.SlaveRouter = MainServerObj.Group("/slaves")

	// Routers
	MainServerObj.GET("/", handlers.Home)
	routers.InitSlavesRouter()
	// Start server
	addr := fmt.Sprintf("%s:%d", src.GlobalSettings.ServerHost, src.GlobalSettings.ServerPort)
	err := MainServerObj.Start(addr)
	if err != nil {
		panic(err)
	}
}
