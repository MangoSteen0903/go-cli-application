package cli_test

import (
	"errors"
	"testing"

	"github.com/MangoSteen0903/go-cli-application/hello-cli/cli"
)

type testValidateArgsConfig struct {
	config cli.Config
	err    error
}

func TestValidateArgs(t *testing.T) {
	tests := []testValidateArgsConfig{
		{
			config: cli.Config{
				NumTimes: -1,
			},
			err: errors.New("must specify number greater than 0"),
		},
		{
			config: cli.Config{},
			err:    errors.New("must specify number greater than 0"),
		},
		{
			config: cli.Config{NumTimes: 10},
			err:    nil,
		},
	}

	for _, tc := range tests {
		err := cli.ValidateArgs(tc.config)

		if tc.err != nil && tc.err.Error() != err.Error() {
			t.Fatalf("Expected error : %v, but got : %v", tc.err, err)
		}

		if tc.err == nil && err != nil {
			t.Fatalf("Expected nil error but got : %v", err)
		}
	}
}
