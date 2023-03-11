package cmd

import (
	"bytes"
	"testing"
)

type testHTMLConfig struct {
	args   []string
	err    error
	output string
}

var usageString = `
http: A HTTP client.
http: <options> server


Options: 
  -method string
    	A method that you want to request. (default "GET")
`

func TestHandleHTML(t *testing.T) {
	tests := []testHTMLConfig{
		{
			args: []string{},
			err:  ErrInvalidNumOfPositionalArgs,
		},
		{
			args:   []string{"-h"},
			err:    nil,
			output: usageString,
		},
		{
			args: []string{"-method", "POST"},
			err:  ErrInvalidNumOfPositionalArgs,
		},
		{
			args: []string{"-method", "POST", "localhost"},
			err:  nil,
		},
	}
	byteBuff := new(bytes.Buffer)
	for _, tc := range tests {
		err := HandleHTML(byteBuff, tc.args)

		if tc.err == nil && err != nil {
			t.Fatalf("Expected nil error, but got : %v\n", err)
		}

		if tc.err != nil && tc.err.Error() != err.Error() {
			t.Fatalf("Expected error : %v, but got : %v\n", tc.err, err)
		}

		output := byteBuff.String()
		if len(tc.output) != 0 {
			if output != tc.output {
				t.Errorf("Expected output : %v\n but got : %v\n", tc.output, output)
			}
		}
		byteBuff.Reset()
	}
}
