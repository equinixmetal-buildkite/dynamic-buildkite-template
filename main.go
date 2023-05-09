package main

import (
	"dynamic-buildkite-template/generator"
	"dynamic-buildkite-template/types"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	trivyPlugin := flag.String("trivyPlugin", "v1.18.2", "provide trivy plugin version")
	shellPlugin := flag.String("shellPlugin", "", "provide shell plugin version")
	ignoreUnfixed := flag.Bool("ignoreUnfixed", true, "provide if unfixed items are to be ignored")
	skipFiles := flag.String("skipFiles", "cosign.key", "provide files to be skipped in trivy plugin")
	shellCheckFiles := flag.String("shellCheckFiles", "script.sh", "provide files to be checked by the shell plugin")

	var severity types.ArrayFlags
	flag.Var(&severity, "severity", "provide the severity")
	var securityChecks types.ArrayFlags
	flag.Var(&securityChecks, "securityChecks", "provide the security checks")

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

	err := generator.GenerateTrivyStep(*trivyPlugin, *shellPlugin, severity, *ignoreUnfixed, securityChecks, *skipFiles, *shellCheckFiles, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
