package cmd

import (
	"errors"
	"flag"
	"fmt"
	"io"
)

type grpcConfig struct {
	server string
	method string
	body   string
}

func HandlegRPC(w io.Writer, args []string) error {
	config := grpcConfig{}
	fs := flag.NewFlagSet("grpc", flag.ContinueOnError)
	fs.SetOutput(w)

	fs.Usage = func() {
		var usageString = `
grpc: A gRPC client.
grpc: <options> server`
		fmt.Fprintln(w, usageString)
		fmt.Fprintln(w)
		fmt.Fprintln(w)
		fmt.Fprintln(w, "Options: ")
		fs.PrintDefaults()
	}
	fs.StringVar(&config.method, "method", "", "A method that you want to request.")
	fs.StringVar(&config.body, "body", "", "grpc Body.")
	err := fs.Parse(args)
	if fs.NArg() != 1 {
		return errors.New("invalid number of positional argument")
	}
	if err != nil {
		return err
	}
	config.server = fs.Arg(0)

	fmt.Fprintln(w, "Running gRPC client.")
	return nil
}
