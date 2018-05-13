package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func environ() map[string]string {
	envMap := make(map[string]string)

	for _, e := range os.Environ() {
		if i := strings.Index(e, "="); i >= 0 {
			envMap[e[:i]] = e[i+1:]
		}
	}

	return envMap
}

func envy(dest string, name string, templates ...string) error {
	tmpls, err := template.New(name).ParseFiles(templates...)
	if err != nil {
		return err
	}

	os.MkdirAll(filepath.Dir(dest), os.ModePerm)
	destWriter, err := os.Create(dest)
	if err != nil {
		return err
	}

	err = tmpls.Execute(destWriter, environ())
	destWriter.Close()

	return err
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s <template>... <dest>\n", os.Args[0])
		flag.PrintDefaults()
	}

	var name string
	flag.StringVar(&name, "name", "", "template to render, defaults to first <template>")

	flag.Parse()

	if len(os.Args) < 3 {
		flag.Usage()
		os.Exit(2)
	}

	args := flag.Args()
	templates, dest := args[:len(args)-1], args[len(args)-1]

	if name == "" {
		name = filepath.Base(templates[0])
	}

	err := envy(dest, name, templates...)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
}
