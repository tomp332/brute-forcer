package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	manager "github.com/tomp332/gobrute/src/ manager"
	"github.com/tomp332/gobrute/src/ manager/api/middlewares"
	"github.com/tomp332/gobrute/src/ manager/api/routers"
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
	routers.BruteRouter = v1ApiGroup.Group("/brute")
	routers.SlaveRouter = v1ApiGroup.Group("/slaves")
	routers.CredsRouter = v1ApiGroup.Group("/creds")

	// Routers
	err := routers.MainRoute{}.InitRouter()
	if err != nil {
		return
	}
	fmt.Printf("Successfully initialized Main Router")
	err = routers.BruteRoute{}.InitRouter()
	if err != nil {
		return
	}
	fmt.Printf("Successfully initialized Brute Router")
	err = routers.CredsRoute{}.InitRouter()
	if err != nil {
		return
	}
	fmt.Printf("Successfully initialized Creds Router")
	err = routers.SlaveRoute{}.InitRouter()
	if err != nil {
		return
	}
	fmt.Printf("Successfully initialized Slave Router")
	// Start server
	addr := fmt.Sprintf("%s:%d", manager.GlobalSettings.ServerHost, manager.GlobalSettings.ServerPort)
	err = MainServerObj.Start(addr)
	if err != nil {
		panic(err)
	}
}
