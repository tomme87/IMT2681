package main

import (
	"bytes"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestProjectInfo_Add(t *testing.T) {
	content, err := ioutil.ReadFile("repos_apache_kafka.json")
	if err != nil {
		t.Error("Unable to read file")
		return
	}
	pi := ProjectInfo{}
	addErr := pi.Add(bytes.NewReader(content)) // https://stackoverflow.com/questions/29746123/convert-byte-array-to-io-read-in-golang
	if addErr != nil {
		t.Error("Unable to add info", addErr)
	}

	if pi.Project != "github.com/apache/kafka" {
		t.Error("Wrong project")
	}
	if pi.Owner != "apache" {
		t.Error("Wrong owner")
	}
}

func TestProjectInfo_AddCommitInfo(t *testing.T) {
	content, err := ioutil.ReadFile("repos_apache_kafka_contributors.json")
	if err != nil {
		t.Error("Unable to read contributors file")
		return
	}
	pi := ProjectInfo{}
	addErr := pi.AddCommitInfo(bytes.NewReader(content)) // https://stackoverflow.com/questions/29746123/convert-byte-array-to-io-read-in-golang
	if addErr != nil {
		t.Error("Unable to add info", addErr)
	}

	if pi.Committer != "ijuma" {
		t.Error("Wrong committer", pi.Committer)
	}
	if pi.Commits != 4176 {
		t.Error("Wrong amount of commits", pi.Commits)
	}
}

func TestProjectInfo_AddLanguageInfo(t *testing.T) {
	content, err := ioutil.ReadFile("repos_apache_kafka_languages.json")
	if err != nil {
		t.Error("Unable to read file")
		return
	}
	pi := ProjectInfo{}
	addErr := pi.AddLanguageInfo(bytes.NewReader(content)) // https://stackoverflow.com/questions/29746123/convert-byte-array-to-io-read-in-golang
	if addErr != nil {
		t.Error("Unable to add info", addErr)
	}

	if !reflect.DeepEqual(pi.Languages, []string{"Java", "Scala", "Python", "Shell", "Batchfile"}) {
		t.Error("Wrong Languages", pi.Languages)
	}
}
