package cmd

import (
	"bytes"
	"testing"
)

type testgRPCConfig struct {
	args   []string
	err    error
	output string
}

var usagegRPCString = `
grpc: A gRPC client.
grpc: <options> server


Options: 
  -body string
    	grpc Body.
  -method string
    	A method that you want to request.
`

func TestHandlegRPC(t *testing.T) {
	tests := []testgRPCConfig{
		{
			args: []string{},
			err:  ErrInvalidNumOfPositionalArgs,
		},
		{
			args:   []string{"-h"},
			err:    nil,
			output: usagegRPCString,
		},
		{
			args: []string{"-method", "POST"},
			err:  ErrInvalidNumOfPositionalArgs,
		},
		{
			args: []string{"-body", "hello"},
			err:  ErrInvalidNumOfPositionalArgs,
		},
		{
			args: []string{"-body", "hello", "localhost"},
			err:  nil,
		},
	}
	byteBuff := new(bytes.Buffer)
	for _, tc := range tests {
		err := HandlegRPC(byteBuff, tc.args)

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
