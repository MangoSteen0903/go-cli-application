package httpClient

import "testing"

type testGetJsonData struct {
	url      string
	filename string
}

func TestGetJsonData(t *testing.T) {
	tests := []testGetJsonData{
		{
			url:      "https://dummyjson.com/quotes",
			filename: "test.json",
		},
	}

	for _, tc := range tests {
		err := GetJsonData(tc.url, tc.filename)
		t.Log(err)
	}
}
