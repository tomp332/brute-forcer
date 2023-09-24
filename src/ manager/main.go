package manager

import (
	_ "github.com/tomp332/gobrute/docs"
	"github.com/tomp332/gobrute/src/ manager/api"
)

// @title Brute Forcer API
// @version 1.0
// @host localhost:3000
// @BasePath /api/v1
func main() {
	// Echo instance
	LoadSettings()
	InitDB()
	api.InitServer()
}
