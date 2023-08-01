package main

import (
	_ "github.com/tomp332/bruteForcer/docs"
	"github.com/tomp332/bruteForcer/src"
	"github.com/tomp332/bruteForcer/src/api"
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
