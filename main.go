package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/tomp332/bruteForcer/src"
	"github.com/tomp332/bruteForcer/src/api"
)

func main() {
	// Echo instance
	src.LoadSettings()
	src.InitDB()
	e := echo.New()
	e.GET("/", api.Home)
	e.POST("/slaves", api.AddSlaves)
	e.GET("/slaves", api.GetSlaves)
	addr := fmt.Sprintf("%s:%d", src.GlobalSettings.ServerHost, src.GlobalSettings.ServerPort)
	err := e.Start(addr)
	if err != nil {
		panic(err)
	}
}
