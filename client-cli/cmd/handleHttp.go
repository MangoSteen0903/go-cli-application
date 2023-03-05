package cmd

import (
	"flag"
	"fmt"
	"io"
)

type httpConfig struct {
	method string
	server string
}

func HandleHttp(w io.Writer, args []string) error {
	var v string
	fs := flag.NewFlagSet("http", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.Usage = func() {
		var usageString = `
http: A HTTP client.
http: <options> server`
		fmt.Fprintln(w, usageString)
		fmt.Fprintln(w)
		fmt.Fprintln(w)
		fmt.Fprintln(w, "Options: ")
		fs.PrintDefaults()

	}

	fs.StringVar(&v, "verb", "GET", "HTTP methods")

	err := fs.Parse(args)

	if err != nil {
		return err
	}

	if v != "GET" && v != "POST" && v != "HEAD" {
		return ErrInvalidMethods
	}

	if fs.NArg() != 1 {
		return ErrNotServerSpecified
	}

	c := httpConfig{method: v}
	c.server = fs.Arg(0)
	fmt.Fprintln(w, "Executing http command")
	return nil
}
