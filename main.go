package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"dynamic-buildkite-template/generator"
)

func main() {
	trivyPlugin := flag.String("trivyPlugin", "v1.18.0", "provide trivy plugin version")
	shellPlugin := flag.String("shellPlugin", "v1.3.0", "provide shell plugin version")

	flag.Usage = func() {
		usage := `
Usage of dynamic-buildkite-template
This Program generates trivy step for the provided options
Options:
`
		fmt.Fprint(os.Stderr, usage)
		flag.PrintDefaults()
	}

	flag.Parse()

	err := generator.GenerateTrivyStep(*trivyPlugin, *shellPlugin, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
