package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func startTestHTTPServer() *httptest.Server {
	pkgResource := `[
		{"name": "package1", "version": "1.1"},
		{"name": "package2", "version": "1.0"}
		]`
	ts := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprint(w, pkgResource)
			},
		),
	)
	return ts
}

func TestFetchRemoteResource(t *testing.T) {
	ts := startTestHTTPServer()
	defer ts.Close()

	data, err := fetchRemoteResource(ts.URL)

	if err != nil {
		t.Fatal(err)
	}

	if len(data) != 2 {
		t.Fatalf("Expected 2 packages, got : %v", data)
	}
}
