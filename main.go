package main

import (
	"dynamic-buildkite-template/generator"
	"os"
)

func main() {
	var trivyPlugin = "v1.18.0"
	var shellPlugin = "v1.3.0"

	generator.GenerateTrivyStep(trivyPlugin, shellPlugin, os.Stdout)
}
