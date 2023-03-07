package main

import (
	"dynamic-buildkite-template/generator"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var trivyPlugin := flag.String("trivyPlugin", "v1.18.0", "provide trivy plugin version")
	var shellPlugin := flag.String("trivyPlugin", "v1.3.0", "provide shell plugin version")

	flag.Usage = func() {
		Usage := `
		Usage of dynamic-buildkite-template
		This Program generates trivy step for the provided options
		Options:
		`
				fmt.Fprintf(os.Stderr, usage)
				flag.PrintDefaults()
	}

	flag.Parse()

	err := generator.GenerateTrivyStep(*trivyPlugin, *shellPlugin, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
