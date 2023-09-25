package main

import (
	_ "github.com/tomp332/gobrute/docs"
	"github.com/tomp332/gobrute/pkg/manager"
	"github.com/tomp332/gobrute/pkg/manager/api"
)

// @title Brute Forcer API
// @version 1.0
// @host localhost:3000
// @BasePath /api/v1
func main() {
	// Echo instance
	manager.LoadManagerSettings()
	manager.InitManagerDB()
	api.InitServer()
}
