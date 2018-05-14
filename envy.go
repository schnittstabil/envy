package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
)

func environ() map[string]interface{} {
	envMap := make(map[string]interface{})

	for _, e := range os.Environ() {
		if i := strings.Index(e, "="); i >= 0 {
			envMap[e[:i]] = e[i+1:]
		}
	}

	return envMap
}

func openOutputWriter(output string) (*os.File, error) {
	if output == "" {
		return os.Stdout, nil
	}

	os.MkdirAll(filepath.Dir(output), os.ModePerm)
	return os.Create(output)
}

func closeOutputWriter(outputWriter *os.File) error {
	if outputWriter == os.Stdout {
		return nil
	}
	return outputWriter.Close()
}

func envy(output string, input string, templates ...string) error {
	tmpls, err := template.New(input).Funcs(sprig.TxtFuncMap()).ParseFiles(templates...)
	if err != nil {
		return err
	}

	outputWriter, err := openOutputWriter(output)
	if err != nil {
		return err
	}

	err = tmpls.Execute(outputWriter, environ())
	closeOutputWriter(outputWriter)

	return err
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [opts] <template>...\n", os.Args[0])
		flag.PrintDefaults()
	}

	var input string
	flag.StringVar(&input, "input", "", "<file> to render, defaults to first <template>")
	var output string
	flag.StringVar(&output, "output", "", "write output to <file> instead of stdout")

	flag.Parse()

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(2)
	}

	templates := flag.Args()

	if input == "" {
		input = filepath.Base(templates[0])
	}

	err := envy(output, input, templates...)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
}
