package cli

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/MangoSteen0903/go-cli-application/cli/util"
)

type Config struct {
	NumTimes     int
	FileLocation string
	IsPrintUsage bool
}

func ParseArgs(w io.Writer, args []string) (Config, error) {
	c := Config{}
	fs := flag.NewFlagSet("hello", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.IntVar(&c.NumTimes, "n", 0, "Number of times to hello")
	fs.StringVar(&c.FileLocation, "o", "", "File Location that you want to save")
	err := fs.Parse(args)

	if err != nil {
		return c, err
	}
	if fs.NArg() != 0 {
		return c, errors.New("positional arguments specified")
	}
	return c, nil
}

func ValidateArgs(c Config) error {
	if !(c.NumTimes > 0) && !c.IsPrintUsage {
		return errors.New("must specify number greater than 0")
	}
	return nil
}

func getName(w io.Writer, r io.Reader) (string, error) {
	msg := "Welcome to Hello Application! Please write your name and hit the enter."
	fmt.Fprintln(w, msg)

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	name := scanner.Text()

	if len(name) == 0 {
		return "", errors.New("you must enter your name. please try again")
	}

	return name, nil
}

func iterateName(w io.Writer, name string, c Config) {
	numTimes := c.NumTimes
	msg := fmt.Sprintf("Hello %s!", name)
	for i := 0; i < numTimes; i++ {
		fmt.Fprintln(w, msg)
	}
}

func saveFile(c Config, name string) error {
	fileName := fmt.Sprintf("%s.html", name)
	fileTemplate := util.GetHTML(name)
	file, err := os.Create(c.FileLocation + "/" + fileName)
	if err != nil {
		return err
	}
	w := bufio.NewWriter(file)
	_, err = w.WriteString(fileTemplate)
	if err != nil {
		return err
	}
	w.Flush()
	return nil
}
func RunApplication(w io.Writer, r io.Reader, c Config) error {

	result, err := getName(w, r)

	if err != nil {
		return err
	}

	err = saveFile(c, result)
	if err != nil {
		return err
	}

	iterateName(w, result, c)

	return nil

}
