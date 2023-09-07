package cmd

import (
	"dynamic-buildkite-template/config"
	"dynamic-buildkite-template/generator"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	g                     = generator.Generator{}
	ConfigFilePath        string
	defaultConfigFilePath = "conf.yaml"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates plugin step for the provided plugins with configurations",
	Long: `
Usage of dynamic-buildkite-template
This Program generates step for the provided plugins with configurations
	`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	generateCmd.PersistentFlags().StringVar(&ConfigFilePath, "config", "", fmt.Sprintf("Mention the config file path (default \"%s\")", defaultConfigFilePath))
}

func initConfig() {
	if ConfigFilePath != "" {
		log.Debug("Config Path:", ConfigFilePath)
		if err := config.LoadConfig(ConfigFilePath); err != nil {
			log.Fatal("Error while loading the configuration file. Exiting the program.")
		}
	} else {
		log.Debug("Config Path:", defaultConfigFilePath)
		if err := config.LoadConfig(defaultConfigFilePath); err != nil {
			log.Debug("Error while loading the configuration file. Loading the defaults")
		}
	}
}

func Execute() error {
	return generateCmd.Execute()
}
