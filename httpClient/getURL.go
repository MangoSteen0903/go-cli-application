package httpClient

import (
	"errors"
	"net/http"
	"strings"
)

var ErrHeaderDoesNotMatch = errors.New("header is not application/json")

func GetURL(url string) (*http.Response, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	header := res.Header["Content-Type"][0]
	headerType := strings.Split(header, ";")[0]
	if headerType != "application/json" {
		return nil, ErrHeaderDoesNotMatch
	}
	return res, err
}
