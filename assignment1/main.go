package main

import (
	"net/http"
	"os"
)

const (
	BasePath = "/projectinfo/v1/"
)

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc(BasePath, handleGetProjectinfo)
	http.ListenAndServe(":"+port, nil)
}
