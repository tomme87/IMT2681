package main

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"
)

// ProjectInfo holds the info about project.
type ProjectInfo struct {
	Project   string   `json:"project"`
	Owner     string   `json:"owner"`
	Committer string   `json:"committer"`
	Commits   int      `json:"commits"`
	Languages []string `json:"languages"`
}

// GitRepo is used to decode repo info into. Could be placed inside of Add().
type GitRepo struct {
	FullName string `json:"full_name"`
	Owner    struct {
		Login string
	}
}

// GitContributor is used to decode contributor info into. Could be placed inside of AddCommitInfo().
type GitContributor struct {
	Login         string
	Contributions int
}

// Add adds project and owner information to projectinfo from JSON
func (pi *ProjectInfo) Add(r io.Reader) error {
	repo := GitRepo{}
	err := json.NewDecoder(r).Decode(&repo)
	if err != nil {
		return err
	}

	pi.Project = fmt.Sprintf("github.com/%s", repo.FullName)
	pi.Owner = repo.Owner.Login

	return nil
}

// AddCommitInfo adds Comitter and commits to projectinfo from JSON
func (pi *ProjectInfo) AddCommitInfo(r io.Reader) error {
	contributors := []GitContributor{}
	err := json.NewDecoder(r).Decode(&contributors)
	if err != nil {
		return err
	}

	pi.Committer = contributors[0].Login
	pi.Commits = contributors[0].Contributions

	return nil
}

// AddLanguageInfo adds languages to projectinfo from JSON
func (pi *ProjectInfo) AddLanguageInfo(r io.Reader) error {
	languages := make(map[string]interface{})
	err := json.NewDecoder(r).Decode(&languages)
	if err != nil {
		return err
	}

	for k := range languages {
		pi.Languages = append(pi.Languages, k)
	}
	sort.Strings(pi.Languages) // Sort alphabetically.

	return nil
}