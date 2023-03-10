package main

import (
	"fmt"
	"io"
	"os"

	"github.com/MangoSteen0903/go-cli-application/cmd"
)

func printUsage(w io.Writer, command string) {
	fmt.Fprintln(w, "Usage: json-downloader [html|grpc] -h")
	switch command {
	case "html":
		cmd.HandleHTML(w, []string{"-h"})
	case "grpc":
		cmd.HandlegRPC(w, []string{"-h"})
	default:
		cmd.HandleHTML(w, []string{"-h"})
		cmd.HandlegRPC(w, []string{"-h"})
	}
}
func handleCommand(w io.Writer, args []string) error {
	var err error
	var cmdType string
	if len(args) < 1 {
		err = cmd.ErrInvalidNumofArgs
	} else {
		switch args[0] {
		case "html":
			err = cmd.HandleHTML(w, args[1:])
			cmdType = "html"
		case "grpc":
			err = cmd.HandlegRPC(w, args[1:])
			cmdType = "grpc"
		case "-h":
			printUsage(w, "")
		case "-help":
			printUsage(w, "")
		default:
			err = cmd.ErrInvalidSubCommand
		}
	}
	if err != nil {
		fmt.Fprintln(w, err)
		printUsage(w, cmdType)
	}
	return err
}
func main() {
	err := handleCommand(os.Stdout, os.Args[1:])
	if err != nil {
		os.Exit(1)
	}
}
