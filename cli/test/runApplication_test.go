package cli_test

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/MangoSteen0903/go-cli-application/cli"
)

type testRunApplicationConfig struct {
	config cli.Config
	input  string
	output string
	err    error
}

func TestRunApplication(t *testing.T) {
	tests := []testRunApplicationConfig{
		{
			config: cli.Config{
				NumTimes: 5,
			},
			input:  "",
			output: strings.Repeat("Welcome to Hello Application! Please write your name and hit the enter.\n", 1),
			err:    errors.New("you must enter your name. please try again"),
		},
		{
			config: cli.Config{
				NumTimes: 5,
			},
			input:  "Milky",
			output: "Welcome to Hello Application! Please write your name and hit the enter.\n" + strings.Repeat("Hello Milky!\n", 5),
			err:    nil,
		},
	}

	for _, tc := range tests {
		bytesBuff := new(bytes.Buffer)
		reader := strings.NewReader(tc.input)
		err := cli.RunApplication(bytesBuff, reader, tc.config)

		if tc.err != nil && tc.err.Error() != err.Error() {
			t.Fatalf("Expected error: %v, but got : %v", tc.err, err)
		}

		if tc.err == nil && err != nil {
			t.Fatalf("Expected nil error but got : %v", err)
		}

		gotMsg := bytesBuff.String()

		if gotMsg != tc.output {
			t.Fatalf("Expected Msg : %v, but got : %v", tc.output, gotMsg)
		}
		bytesBuff.Reset()
	}
}
