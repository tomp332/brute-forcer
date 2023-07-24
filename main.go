package main

import (
	"github.com/tomp332/bruteForcer/src"
	"github.com/tomp332/bruteForcer/src/api"
)

func main() {
	// Echo instance
	src.LoadSettings()
	src.InitDB()
	api.InitServer()
}
