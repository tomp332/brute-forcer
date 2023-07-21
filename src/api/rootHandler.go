package api

import (
	"fmt"
	"io"
	"net/http"
)

// GetRoot Create home HTTP root
func GetRoot(w http.ResponseWriter, _ *http.Request) {
	fmt.Printf("Welcome to the home page!\n")
	_, err := io.WriteString(w, "Hello from a GetRoot #1!\n")
	if err != nil {
		return
	}
}
