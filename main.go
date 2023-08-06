package main

import (
	_ "github.com/tomp332/gospray/docs"
	"github.com/tomp332/gospray/src"
	"github.com/tomp332/gospray/src/api"
)

// @title Brute Forcer API
// @version 1.0
// @host localhost:3000
// @BasePath /api/v1
func main() {
	// Echo instance
	src.LoadSettings()
	src.InitDB()
	api.InitServer()
}
