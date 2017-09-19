package main

import (
	"net/http"
)

const (
	BasePath = "/projectinfo/v1/"
)

func main() {
	http.HandleFunc(BasePath, handleGetProjectinfo)
	http.ListenAndServe(":8080", nil)
}
