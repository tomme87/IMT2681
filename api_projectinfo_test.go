package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
 Testing that we get not implemented when doing something other than a GET request.

 The GET request is not tested since we do not want to query GitHub in our test.
 Most of this is tested in projectinfo_test.go anyway.
*/
func TestHandleGetProjectinfo_NotImplemented(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(handleGetProjectinfo))
	defer ts.Close()

	client := &http.Client{}

	testMethods := []string{http.MethodDelete, http.MethodPost, http.MethodPut}

	for _, method := range testMethods {
		req, err := http.NewRequest(method, ts.URL, nil)
		if err != nil {
			t.Errorf("Error constructing %s request. Error: %s", method, err)
			continue
		}

		resp, err := client.Do(req)
		if err != nil {
			t.Errorf("Error executing %s request. Error %s", method, err)
			continue
		}

		if resp.StatusCode != http.StatusNotImplemented {
			t.Errorf("Expected %s, but got %s", http.StatusNotImplemented, resp.StatusCode)
		}
	}
}
