package cmd

import (
	"flag"
	"fmt"
	"io"

	"github.com/MangoSteen0903/go-cli-application/httpClient"
)

type htmlConfig struct {
	url      string
	method   string
	filename string
}

func HandleHTML(w io.Writer, args []string) error {
	config := htmlConfig{}
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
	fs.StringVar(&config.method, "method", "GET", "A method that you want to request.")
	fs.StringVar(&config.filename, "filename", "test.json", "A Filename that you want to save it.")
	err := fs.Parse(args)
	if err != nil {
		switch {
		case err.Error() == ErrHelpRequest.Error():
			return nil
		default:
			return err
		}
	}

	if fs.NArg() != 1 {
		return ErrInvalidNumOfPositionalArgs
	}
	config.url = fs.Arg(0)
	err = httpClient.GetJsonData(config.url, config.filename)
	if err != nil {
		return err
	}
	return nil
}
