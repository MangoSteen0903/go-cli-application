package htmlClient

import (
	"errors"
	"testing"
)

type testGetURLConfig struct {
	url string
	err error
}

func TestGetURL(t *testing.T) {
	tests := []testGetURLConfig{
		{
			url: "asdf",
			err: errors.New(`Get "asdf": unsupported protocol scheme ""`),
		},
		{
			url: "https://naver.com",
			err: ErrHeaderDoesNotMatch,
		},
		{
			url: "https://dummyjson.com/quotes",
			err: nil,
		},
	}

	for _, tc := range tests {
		_, err := GetURL(tc.url)
		if tc.err == nil && err != nil {
			t.Fatalf("Expected nil error but got : %v", err)
		}
		if tc.err != nil && tc.err.Error() != err.Error() {
			t.Fatalf("Expected error : %v, but got : %v", tc.err, err)
		}
	}
}
