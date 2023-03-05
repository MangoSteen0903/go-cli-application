package cli_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/MangoSteen0903/go-cli-application/hello-cli/cli"
)

type testParseArgsConfig struct {
	args   []string
	err    error
	output string
	config cli.Config
}

func TestParseArgs(t *testing.T) {
	test := []testParseArgsConfig{
		{
			args: []string{"-h"},
			err:  errors.New("flag: help requested"),
			output: `
A Hello Application which prints the name you entered a specified number of times.

Usage of hello: <options> [name]
Options:
  -n int
	Number of times that iterate your name.
  -o string
	Location that you want to save the file.`,
			config: cli.Config{NumTimes: 0},
		},
		{
			args: []string{"-help"},
			err:  errors.New("flag: help requested"),
			output: `
A Hello Application which prints the name you entered a specified number of times.

Usage of hello: <options> [name]
Options:
  -n int
	Number of times that iterate your name.
  -o string
	Location that you want to save the file.`,
			config: cli.Config{NumTimes: 0},
		},
	}
	byteBuff := new(bytes.Buffer)
	for _, tc := range test {
		c, err := cli.ParseArgs(byteBuff, tc.args)

		if tc.err != nil && err.Error() != tc.err.Error() {
			t.Fatalf("Expected Err: %v, But got: %v\n", tc.err, err)
		}

		if tc.err == nil && err != nil {
			t.Fatalf("Expected nil err but got: %v\n", err)
		}
		gotMsg := byteBuff.String()
		if c.NumTimes != tc.config.NumTimes {
			t.Fatalf("Expected Num times %v, but got : %v", tc.config.NumTimes, c.NumTimes)
		}

		if len(tc.output) != 0 && gotMsg != tc.output {
			t.Fatalf("Expected output : %v, but got: %v", gotMsg, tc.output)
		}

		byteBuff.Reset()
	}
}
