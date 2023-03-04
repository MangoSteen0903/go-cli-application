package cli

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Config struct {
	NumTimes     int
	IsPrintUsage bool
}

var PRINT_USAGE = fmt.Sprintf(`Usage: %s <interger> [-h|-help]

A Hello Application which prints the name you entered <integer> number of times.
`, os.Args[0])

func PrintUsage(w io.Writer) {
	fmt.Fprint(w, PRINT_USAGE)
}

func ParseArgs(args []string) (Config, error) {
	var numTimes int
	var err error
	c := Config{}

	if len(args) != 1 {
		return c, errors.New("invalid number of arguments")
	}

	if args[0] == "-h" || args[0] == "-help" {
		c.IsPrintUsage = true
		return c, nil
	}
	numTimes, err = strconv.Atoi(args[0])
	if err != nil {
		return c, err
	}
	c.NumTimes = numTimes
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

func RunApplication(w io.Writer, r io.Reader, c Config) error {

	if c.IsPrintUsage {
		PrintUsage(w)
		return nil
	}
	result, err := getName(w, r)

	if err != nil {
		return err
	}

	iterateName(w, result, c)

	return nil

}
