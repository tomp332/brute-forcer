package main

import (
	"github.com/tomp332/bruteForcer/src/api"
	"net/http"
)

func setupHandlers() {
	http.HandleFunc("/api", api.GetRoot)
}

func main() {
	// Echo instance
	//mainSettings := src.LoadSettings()
	//e := echo.New()

}
