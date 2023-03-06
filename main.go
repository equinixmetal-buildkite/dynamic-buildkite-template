package main

import (
	"dynamic-buildkite-template/buildkite"
	//"fmt"
)

func main() {
	var trivyPlugin = "v1.18.0"
	var shellPlugin = "v1.3.0"

	buildkite.TrivyStepGenerator(trivyPlugin, shellPlugin)

	//fmt.Println("Hello World")
}
