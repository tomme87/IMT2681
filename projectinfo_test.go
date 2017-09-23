package main

import (
	"bytes"
	"io/ioutil"
	"reflect"
	"testing"
)

/**
 * Testing Add() in projectinfo.
 * Using the file repos_apache_kafka.json file, instead of querying github for live data.
 */
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

	if pi.Project != "kafka" {
		t.Error("Wrong project", pi.Project)
	}
	if pi.Owner != "apache" {
		t.Error("Wrong owner")
	}
}

/**
 * Testing AddCommitInfo() in projectinfo
 * Using the file repos_apache_kafka_contributors.json file, instead of querying github for live data.
 */
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
	if pi.Commits != 309 {
		t.Error("Wrong amount of commits", pi.Commits)
	}
}

/**
 * Testing AddLanguageInfo() in projectinfo
 * Using the file repos_apache_kafka_languages.json file, instead of querying github for live data.
 */
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

	if !reflect.DeepEqual(pi.Languages, []string{"Batchfile", "HTML", "Java", "Python", "Scala", "Shell", "XSLT"}) {
		t.Error("Wrong Languages", pi.Languages)
	}
}
