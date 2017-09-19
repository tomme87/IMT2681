package main

import (
	"net/http"
	"strings"
	"fmt"
)

const (
	Host  = 3
	Owner = 4
	Repo  = 5
)

func handleGetProjectinfo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Add("content-type", "application/json")
		parts := strings.Split(r.URL.Path, "/") // 3 = HOST, 4 = owner, 5 = repo
		if parts[Host] == "github.com" {
			return
		}
		fmt.Fprintf(w, "Hei")

	default:
		http.Error(w, "Not implemented", http.StatusNotImplemented)
		return
	}
}
