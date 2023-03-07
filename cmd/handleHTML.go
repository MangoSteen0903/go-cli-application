package cmd

import (
	"flag"
	"fmt"
	"io"
)

type htmlConfig struct {
	method string
	url    string
}

func HandleHTML(w io.Writer, args []string) error {
	var config htmlConfig
	fs := flag.NewFlagSet("html", flag.ContinueOnError)
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

	fs.StringVar(&config.url, "u", "", "A url that you want to retreive data")
	fs.StringVar(&config.method, "method", "GET", "A method that you want to request on your url.")

	err := fs.Parse(args)

	if err != nil {
		return err
	}

	if fs.NArg() != 1 {
		return ErrNotURLSpecified
	}

	fmt.Fprintln(w, "Running HTML Command")

	return nil
}
