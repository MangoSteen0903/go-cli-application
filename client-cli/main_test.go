package main

import (
	"bytes"
	"testing"

	"github.com/MangoSteen0903/go-cli-application/client-cli/cmd"
)

type testConfig struct {
	args   []string
	err    error
	output string
}

var usageString = `Usage: mync [http|grpc] -h

http: A HTTP client.
http: <options> server


Options: 
  -verb string
        HTTP methods (default "GET")

grpc: A gRPC client.
grpc: <options> server


Options: 
  -body string
        gRPC body
  -method string
        gRPC method that you want to call.
`

func TestMain(t *testing.T) {
	tests := []testConfig{
		{
			args:   []string{},
			err:    cmd.ErrInvalidSubCommand,
			output: "invalid sub-command\n" + usageString,
		},
	}
	byteBuff := new(bytes.Buffer)
	for _, tc := range tests {
		err := handleCommand(byteBuff, tc.args)
		if tc.err == nil && err != nil {
			t.Fatalf("Expected nil error, but got : %v", err)
		}

		if tc.err != nil && tc.err.Error() != err.Error() {
			t.Fatalf("Expected Error %v, but got : %v", tc.err, err)
		}
		output := byteBuff.String()
		if len(tc.output) > 0 {
			if tc.output != output {
				t.Errorf("Expected message :%v, but got :%v", tc.output, output)
			}
		}
	}
}
