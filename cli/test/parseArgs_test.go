package cli_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/MangoSteen0903/go-cli-application/cli"
)

type testParseArgsConfig struct {
	args     []string
	err      error
	numTimes int
}

func TestParseArgs(t *testing.T) {
	test := []testParseArgsConfig{
		{
			args:     []string{"-h"},
			err:      errors.New("flag: help requested"),
			numTimes: 0,
		},
		{
			args:     []string{"-help"},
			err:      errors.New("flag: help requested"),
			numTimes: 0,
		},
		{
			args:     []string{"-n", "3"},
			err:      nil,
			numTimes: 3,
		},
		{
			args:     []string{"-n", "abc"},
			err:      errors.New(`invalid value "abc" for flag -n: parse error`),
			numTimes: 0,
		},
		{
			args:     []string{"-n", "3", "foo"},
			err:      errors.New("positional arguments specified"),
			numTimes: 3,
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

		if c.NumTimes != tc.numTimes {
			t.Fatalf("Expected Num times %v, but got : %v", tc.numTimes, c.NumTimes)
		}

		byteBuff.Reset()
	}
}
