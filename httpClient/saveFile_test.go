package htmlClient

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func startTestHttpSever() *httptest.Server {
	pkgData := `
	{"quotes":[{"id":1,"quote":"Life isn’t about getting and having, it’s about giving and being.","author":"Kevin Kruse"},{"id":2,"quote":"Whatever the mind of man can conceive and believe, it can achieve.","author":"Napoleon Hill"}],"total":100,"skip":0,"limit":2}`
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, pkgData)
		}),
	)
	return ts
}

func TestSaveFile(t *testing.T) {
	ts := startTestHttpSever()
	defer ts.Close()
	res, _ := GetURL(ts.URL)
	pkg, err := SaveFile("test.json", res)
	if err != nil {
		t.Fatal(err)
	}
	if len(pkg.Quotes) != 2 {
		t.Fatalf("Expected 2 packages, got : %v", pkg.Quotes)
	}
	os.Remove("test.json")
}
