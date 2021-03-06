package main

import (
	"net/http"
	"os"
)

// BasePath is the base URL for out project.
const (
	BasePath = "/projectinfo/v1/"
)

/**
 * Main function that set up the HTTP server.
 */
func main() {
	port := os.Getenv("PORT") // Get port from environment variable. Needed to deploy on heruko.
	if port == "" {
		port = "8080" // Default to port 8080
	}
	http.HandleFunc(BasePath, handleGetProjectinfo)
	http.ListenAndServe(":"+port, nil)
}
