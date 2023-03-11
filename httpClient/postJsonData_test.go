package httpClient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func packageRegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		p := QuotesJson{}
		d := PkgRegisterResult{}
		defer r.Body.Close()
		data, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = json.Unmarshal(data, &p)
		if err != nil || len(p.Quote) == 0 {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		d.Id = p.Id + 1
		jsonData, err := json.Marshal(d)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, string(jsonData))
	} else {
		http.Error(w, "Invalid HTTP method Specified", http.StatusMethodNotAllowed)
		return
	}
}

func startTestPackageServer() *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(packageRegisterHandler))
	return ts
}

func TestPostJsonData(t *testing.T) {
	ts := startTestPackageServer()
	defer ts.Close()
	p := QuotesJson{}
	resp, err := RegisterPackageData(ts.URL, p)
	if err == nil {
		t.Fatal("Expected error to be non-nil, got nil")
	}
	if resp.Id != 0 {
		t.Errorf("Expected package ID to be empty, got : %d", resp.Id)
	}
}
