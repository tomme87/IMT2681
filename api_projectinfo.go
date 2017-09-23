package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// 3 = HOST, 4 = owner, 5 = repo
const (
	Host  = 3
	Owner = 4
	Repo  = 5
)

/**
 * Our function that handles our HTTP call to our server.
 *
 * @param w The response writer.
 * @param r The HTTP request.
 */
func handleGetProjectinfo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Add("content-type", "application/json")
		parts := strings.Split(r.URL.Path, "/") // 3 = HOST, 4 = owner, 5 = repo

		//if parts[Host] != "github.com" || parts[Owner] != "" || parts[Repo] != "" {
		if len(parts) < 6 {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		if parts[Host] != "github.com" {
			http.Error(w, "Not implemented. only github.com supported.", http.StatusNotImplemented)
			return
		}

		pi := ProjectInfo{}

		//repoRes, err := http.Get("https://api.github.com/repos/apache/kafka")
		repoRes, err := http.Get(fmt.Sprintf("https://api.github.com/repos/%s/%s", parts[Owner], parts[Repo]))
		if err != nil {
			http.Error(w, "unable to get repo info", http.StatusInternalServerError)
			return
		}
		pi.Add(repoRes.Body)

		contributorsRes, err := http.Get(fmt.Sprintf("https://api.github.com/repos/%s/%s/contributors", parts[Owner], parts[Repo]))
		if err != nil {
			http.Error(w, "unable to get repo info", http.StatusInternalServerError)
			return
		}
		pi.AddCommitInfo(contributorsRes.Body)

		languagesRes, err := http.Get(fmt.Sprintf("https://api.github.com/repos/%s/%s/languages", parts[Owner], parts[Repo]))
		if err != nil {
			http.Error(w, "unable to get repo info", http.StatusInternalServerError)
			return
		}
		pi.AddLanguageInfo(languagesRes.Body)

		json.NewEncoder(w).Encode(pi)

	default:
		http.Error(w, "Not implemented", http.StatusNotImplemented)
		return
	}
}
