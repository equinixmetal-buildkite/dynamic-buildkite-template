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

	// Output to default stderr instead of the stdout
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stderr)

	// Only log the error severity or above.
	log.SetLevel(log.ErrorLevel)

	if err := config.LoadConfig("resources/config"); err != nil {
		panic(err)
	}
}
func main() {
	cmd.Execute()
}
