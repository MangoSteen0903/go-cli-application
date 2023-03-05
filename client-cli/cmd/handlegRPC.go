package cmd

import (
	"flag"
	"fmt"
	"io"
)

type grpcConfig struct {
	method string
	body   string
	server string
}

func HandlegRPC(w io.Writer, args []string) error {
	c := grpcConfig{}
	fs := flag.NewFlagSet("grpc", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.Usage = func() {}

	fs.StringVar(&c.method, "method", "", "gRPC method that you want to call.")
	fs.StringVar(&c.body, "body", "", "gRPC body")
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
	err := fs.Parse(args)

	if err != nil {
		return err
	}

	if fs.NArg() != 1 {
		return ErrNotServerSpecified
	}

	c.server = fs.Arg(0)
	fmt.Fprintln(w, "Executing grpc command")
	return nil
}
