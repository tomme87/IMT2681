package main

import (
	"net/http"
	"os"
)

const (
	BasePath = "/projectinfo/v1/"
)

/**
 * Main function that set up the HTTP server.
 */
func main() {
	port := os.Getenv("PORT") // Get port from environment variable. Needed to deploy on heruko.
	http.HandleFunc(BasePath, handleGetProjectinfo)
	http.ListenAndServe(":"+port, nil)
}
