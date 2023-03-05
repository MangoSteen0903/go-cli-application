package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/MangoSteen0903/go-cli-application/client-cli/cmd"
)

func printUsage(w io.Writer) {
	fmt.Fprintf(w, "Usage: mync [http|grpc] -h\n")
	cmd.HandleHttp(w, []string{"-h"})
	cmd.HandlegRPC(w, []string{"-h"})
}

func handleCommand(w io.Writer, args []string) error {
	var err error

	if len(args) < 1 {
		err = cmd.ErrInvalidSubCommand
	} else {
		switch args[0] {
		case "http":
			err = cmd.HandleHttp(w, args[1:])
		case "grpc":
			err = cmd.HandlegRPC(w, args[1:])
		case "-h":
			printUsage(w)
		case "-help":
			printUsage(w)
		default:
			err = cmd.ErrInvalidSubCommand
		}
	}

	if errors.Is(err, cmd.ErrInvalidSubCommand) || errors.Is(err, cmd.ErrNotServerSpecified) || errors.Is(err, cmd.ErrInvalidMethods) {
		fmt.Fprintln(w, err)
		printUsage(w)
	}
	return err
}
func main() {
	err := handleCommand(os.Stdout, os.Args[1:])
	if err != nil {
		os.Exit(1)
	}
}
