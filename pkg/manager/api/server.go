package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/tomp332/gobrute/pkg/manager"
	"github.com/tomp332/gobrute/pkg/manager/api/middlewares"
	routers2 "github.com/tomp332/gobrute/pkg/manager/api/routers"
)

var MainServerObj *echo.Echo

func InitServer() {
	MainServerObj = echo.New()
	// Middlewares
	MainServerObj.Use(middlewares.LogRequest)

	// Groups
	MainServerObj.GET("/swagger/*", echoSwagger.WrapHandler)
	v1ApiGroup := MainServerObj.Group("/api/v1")
	routers2.MainRouter = v1ApiGroup.Group("")
	routers2.BruteRouter = v1ApiGroup.Group("/brute")
	routers2.SlaveRouter = v1ApiGroup.Group("/slaves")
	routers2.CredsRouter = v1ApiGroup.Group("/creds")

	// Routers
	err := routers2.MainRoute{}.InitRouter()
	if err != nil {
		return
	}
	fmt.Printf("Successfully initialized Main Router")
	err = routers2.BruteRoute{}.InitRouter()
	if err != nil {
		return
	}
	fmt.Printf("Successfully initialized Brute Router")
	err = routers2.CredsRoute{}.InitRouter()
	if err != nil {
		return
	}
	fmt.Printf("Successfully initialized Creds Router")
	err = routers2.SlaveRoute{}.InitRouter()
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
