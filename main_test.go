package main

import (
	"bytes"
	"testing"

	"github.com/MangoSteen0903/go-cli-application/cmd"
)

var usageString = `Usage: json-downloader [html|grpc] -h

http: A HTTP client.
http: <options> server


Options: 
  -method string
        A method that you want to request. (default "GET")

grpc: A gRPC client.
grpc: <options> server


Options: 
  -body string
        grpc Body.
  -method string
        A method that you want to request.
`

type mainConfig struct {
	args   []string
	err    error
	output string
}

func TestXxx(t *testing.T) {
	tests := []mainConfig{
		{
			args:   []string{},
			err:    cmd.ErrInvalidNumofArgs,
			output: "invalid number of arguments\n" + usageString,
		},
	}
	byteBuff := new(bytes.Buffer)

	for _, tc := range tests {
		err := handleCommand(byteBuff, tc.args)

		if tc.err == nil && err != nil {
			t.Fatalf("Expected nil error, but got : %v\n", err)
		}

		if tc.err != nil && tc.err.Error() != err.Error() {
			t.Fatalf("Expected Error: %v, but got : %v\n", tc.err, err)
		}

		output := byteBuff.String()
		if tc.output != "" && tc.output != output {
			t.Fatalf("Expected output: %v, but got: %v", tc.output, output)
		}
		byteBuff.Reset()
	}
}
