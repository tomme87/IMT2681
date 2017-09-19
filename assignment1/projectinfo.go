package main

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"
)

type ProjectInfo struct {
	Project   string
	Owner     string
	Committer string
	Commits   int
	Languages []string
}

type GitRepo struct {
	Full_name string
	Owner     struct {
		Login string
	}
}

type GitContributor struct {
	Login         string
	Contributions int
}

func (pi *ProjectInfo) Add(r io.Reader) error {
	var repo GitRepo
	err := json.NewDecoder(r).Decode(&repo)
	if err != nil {
		return err
	}

	pi.Project = fmt.Sprintf("github.com/%s", repo.Full_name)
	pi.Owner = repo.Owner.Login

	return nil
}

func (pi *ProjectInfo) AddCommitInfo(r io.Reader) error {
	var contributors []GitContributor
	err := json.NewDecoder(r).Decode(&contributors)
	if err != nil {
		return err
	}

	pi.Committer = contributors[0].Login
	pi.Commits = contributors[0].Contributions

	return nil
}

func (pi *ProjectInfo) AddLanguageInfo(r io.Reader) error {
	//var languages map[string]interface{}
	languages := make(map[string]interface{})
	err := json.NewDecoder(r).Decode(&languages)
	if err != nil {
		return err
	}

	pi.Languages = make([]string, len(languages))
	i := 0
	for k := range languages {
		pi.Languages[i] = k
		i++
	}
	sort.Strings(pi.Languages)

	return nil
}

/*
func (pi *ProjectInfo) toJSON(w *io.Writer) {
	json.NewEncoder(w).Encode(pi)
}
*/
