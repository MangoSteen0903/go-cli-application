package httpClient

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func RegisterPackageData(url string, data QuotesJson) (PkgRegisterResult, error) {
	result := PkgRegisterResult{}
	b, err := json.Marshal(data)
	if err != nil {
		return result, err
	}

	reader := bytes.NewReader(b)
	r, err := http.Post(url, "application/json", reader)
	if err != nil {
		return result, err
	}

	defer r.Body.Close()

	respData, err := io.ReadAll(r.Body)
	if err != nil {
		return result, err
	}

	if r.StatusCode != http.StatusOK {
		return result, errors.New(string(respData))
	}
	err = json.Unmarshal(respData, &result)
	return result, err
}
