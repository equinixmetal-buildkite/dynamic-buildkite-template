package main

import (
	"dynamic-buildkite-template/cmd"
	"dynamic-buildkite-template/config"
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the error severity or above.
	log.SetLevel(log.ErrorLevel)

	if err := config.LoadConfig("resources/config"); err != nil {
		panic(err)
	}
}
func main() {
	cmd.Execute()
}
