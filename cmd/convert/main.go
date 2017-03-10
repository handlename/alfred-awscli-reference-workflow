package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/handlename/alfred-awscli-reference-workflow"
)

func main() {
	var srcPath string
	flag.StringVar(&srcPath, "src-path", "", "path to sitemap.xml")
	flag.Parse()

	src, err := os.Open(srcPath)
	if err != nil {
		fmt.Printf("failed to open sitemap.xml: %s", err.Error())
		os.Exit(1)
	}

	if err = wf.Convert(src, os.Stdout); err != nil {
		fmt.Printf("failed to convert sitemap.xml: %s", err.Error())
		os.Exit(1)
	}
}
