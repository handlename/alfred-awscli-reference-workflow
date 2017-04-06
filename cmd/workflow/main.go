package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	wf "github.com/handlename/alfred-awscli-reference-workflow"
)

func main() {
	var (
		path     string
		keywords string
	)

	flag.StringVar(&path, "path", "", "path to commnd list file")
	flag.StringVar(&keywords, "keywords", "", "space separated keyword list")
	flag.Parse()

	if path == "" {
		fmt.Fprintf(os.Stderr, "--path is required\n")
		os.Exit(1)
	}

	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open command list file: %s\n", err.Error())
		os.Exit(1)
	}

	if err = wf.Search(file, strings.Split(keywords, " "), os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "failed to search command: %s\n", err.Error())
		os.Exit(1)
	}
}
