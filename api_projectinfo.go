package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// 3 = HOST, 4 = owner, 5 = repo, BaseURL the base url to github API
const (
	Host    = 3
	Owner   = 4
	Repo    = 5
	BaseURL = "https://api.github.com/"
)

// RateLimit holds the rate limit info from github.
type RateLimit struct {
	Rate struct {
		Limit     int
		Remaining int
		Reset     int
	}
}

// checkRateLimit checks the github rate limit.
func checkRateLimit() (RateLimit, error) {
	res, err := http.Get(BaseURL + "rate_limit")
	if err != nil {
		return RateLimit{}, err
	}
	rl := RateLimit{}
	errJ := json.NewDecoder(res.Body).Decode(&rl)
	if errJ != nil {
		return RateLimit{}, err
	}
	return rl, nil
}

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

		rl, err := checkRateLimit()
		if err != nil { // Unable to check rate limit..
			http.Error(w, "Unable to check rate limit", http.StatusInternalServerError)
			return
		}
		if rl.Rate.Limit == rl.Rate.Remaining { // Rate limit exceeded.
			json.NewEncoder(w).Encode(rl)
			return
		}

		pi := ProjectInfo{}

		//repoRes, err := http.Get("https://api.github.com/repos/apache/kafka")
		repoRes, err := http.Get(fmt.Sprintf("%srepos/%s/%s", BaseURL, parts[Owner], parts[Repo]))
		if err != nil {
			http.Error(w, "unable to get repo info", http.StatusInternalServerError)
			return
		}
		pi.Add(repoRes.Body)

		contributorsRes, err := http.Get(fmt.Sprintf("%srepos/%s/%s/contributors", BaseURL, parts[Owner], parts[Repo]))
		if err != nil {
			http.Error(w, "unable to get repo info", http.StatusInternalServerError)
			return
		}
		pi.AddCommitInfo(contributorsRes.Body)

		languagesRes, err := http.Get(fmt.Sprintf("%srepos/%s/%s/languages", BaseURL, parts[Owner], parts[Repo]))
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
