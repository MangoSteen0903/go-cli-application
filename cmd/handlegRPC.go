package cmd

import (
	"flag"
	"fmt"
	"io"
)

type grpcConfig struct {
	methods string
	body    string
	server  string
}

func HandlegRPC(w io.Writer, args []string) error {
	var config grpcConfig
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

	fs.StringVar(&config.methods, "method", "", "gRPC method that you want to call.")
	fs.StringVar(&config.body, "body", "", "gRPC Body")

	err := fs.Parse(args)

	if err != nil {
		return err
	}

	if fs.NArg() != 1 {
		return ErrNotURLSpecified
	}
	config.server = fs.Arg(0)
	fmt.Fprintln(w, "Running gRPC Command")

	return nil
}
