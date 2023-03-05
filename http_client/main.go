package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type pkgData struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func fetchRemoteResource(url string) ([]pkgData, error) {

	var packages []pkgData
	r, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	if r.Header.Get("Content-Type") != "application/json" {
		return packages, nil
	}

	data, err := io.ReadAll(r.Body)

	if err != nil {
		return packages, err
	}
	err = json.Unmarshal(data, &packages)
	return packages, err
}
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stdout, "Must Specify a HTTP URL to get data from")
		os.Exit(1)
	}

	body, err := fetchRemoteResource(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stdout, "%v\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "%s\n", body)

}
