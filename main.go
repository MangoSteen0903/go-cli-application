package main

import (
	"fmt"
	"io"
	"os"

	"github.com/MangoSteen0903/go-cli-application/cmd"
)

func printUsage(w io.Writer) {
	fmt.Fprintf(w, "Usage: json-downloader [http|grpc] -h\n")
	cmd.HandleHTML(w, []string{"-h"})
	cmd.HandlegRPC(w, []string{"-h"})
}
func handleCommand(w io.Writer, args []string) error {

	var err error

	if len(args) < 1 {
		err = cmd.ErrInvalidArguments
	} else {
		switch args[0] {
		case "html":
			err = cmd.HandleHTML(w, args[1:])
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

	if err == cmd.ErrInvalidArguments || err == cmd.ErrInvalidSubCommand || err == cmd.ErrNotURLSpecified {
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
