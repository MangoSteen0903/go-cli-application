package cli_test

import (
	"errors"
	"testing"

	"github.com/MangoSteen0903/go-cli-application/cli"
)

type testParseArgsConfig struct {
	args   []string
	config cli.Config
	err    error
}

func TestParseArgs(t *testing.T) {
	test := []testParseArgsConfig{
		{
			args: []string{},
			err:  errors.New("invalid number of arguments"),
		},
		{
			args:   []string{"-h"},
			err:    nil,
			config: cli.Config{IsPrintUsage: true},
		},
		{
			args:   []string{"-help"},
			err:    nil,
			config: cli.Config{IsPrintUsage: true},
		},
		{
			args:   []string{"3"},
			err:    nil,
			config: cli.Config{NumTimes: 3},
		},
		{
			args:   []string{"abc"},
			err:    errors.New(`strconv.Atoi: parsing "abc": invalid syntax`),
			config: cli.Config{},
		},
	}

	for _, tc := range test {
		c, err := cli.ParseArgs(tc.args)

		if tc.err != nil && err.Error() != tc.err.Error() {
			t.Fatalf("Expected Err: %v, But got: %v\n", tc.err, err)
		}

		if tc.err == nil && err != nil {
			t.Fatalf("Expected nil err but got: %v\n", err)
		}

		if c.IsPrintUsage != tc.config.IsPrintUsage {
			t.Fatalf("Expected Print Usage %v, but got : %v", tc.config.IsPrintUsage, c.IsPrintUsage)
		}

		if c.NumTimes != tc.config.NumTimes {
			t.Fatalf("Expected Num times %v, but got : %v", tc.config.NumTimes, c.NumTimes)
		}
	}
}
