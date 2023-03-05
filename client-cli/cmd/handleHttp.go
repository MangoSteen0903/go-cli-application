package cmd

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type httpConfig struct {
	method   string
	server   string
	fileName string
}

type pkgData struct {
	Id        int    `json:"id"`
	Todo      string `json:"todo"`
	Completed bool   `json:"completed"`
	UserId    int    `json:"userId"`
}

type entireJsonData struct {
	Todos []pkgData `json:"todos"`
	Total int       `json:"total"`
	Skip  int       `json:"skip"`
	Limit int       `json:"limit"`
}

func HandleHttp(w io.Writer, args []string) error {
	var v string
	var f string
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
	fs.StringVar(&f, "o", "", "File name")
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
	c.fileName = f
	switch c.method {
	case "GET":
		fmt.Println("Execute Get methods")
		pkgs, err := executeGet(c.server)
		if err != nil {
			return err
		}
		err = saveFile(pkgs, c.fileName)
		if err != nil {
			return err
		}
	}
	fmt.Fprintln(w, "Executing http command")
	return nil
}

func executeGet(url string) (entireJsonData, error) {
	pkgs, err := fetchRemoteResource(url)
	if err != nil {
		return pkgs, err
	}
	return pkgs, err
}

func fetchRemoteResource(url string) (entireJsonData, error) {

	var pkgs entireJsonData
	r, err := http.Get(url)
	if err != nil {
		return pkgs, err
	}

	defer r.Body.Close()
	fmt.Println(r.Header.Get("Content-Type"))
	if r.Header.Get("Content-Type") != "application/json; charset=utf-8" {
		return pkgs, nil
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		return pkgs, err
	}
	err = json.Unmarshal(data, &pkgs)
	return pkgs, err
}

func saveFile(pkgs entireJsonData, fileName string) error {
	file, err := os.Create("../files/" + fileName)

	if err != nil {
		return err
	}
	w := bufio.NewWriter(file)
	var todos []string
	for _, pk := range pkgs.Todos {
		todo := fmt.Sprintf(`{"id": %v,"todo": "%v","completed": %v,"userId": %v},`, pk.Id, pk.Todo, pk.Completed, pk.UserId)
		todos = append(todos, todo)
	}
	fmt.Println(todos)
	_, err = w.WriteString(fmt.Sprintf("[%v]", strings.Join(todos, "")))
	if err != nil {
		return err
	}
	w.Flush()
	return nil
}
