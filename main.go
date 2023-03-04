package main

import (
	"fmt"
	"os"

	"github.com/MangoSteen0903/go-cli-application/cli"
)

func main() {
	config, err := cli.ParseArgs(os.Stdout, os.Args[1:])

	if err != nil {
		fmt.Fprintln(os.Stdout, err)

		os.Exit(1)
	}
	err = cli.ValidateArgs(config)

	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}

	err = cli.RunApplication(os.Stdout, os.Stdin, config)

	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}

}
