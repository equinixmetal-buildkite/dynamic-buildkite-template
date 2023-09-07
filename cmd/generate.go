package cmd

import (
	"dynamic-buildkite-template/config"
	"dynamic-buildkite-template/generator"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	g              = generator.Generator{}
	ConfigFilePath string
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
	generateCmd.PersistentFlags().StringVar(&ConfigFilePath, "config", "conf.yaml", "Mention the config file path")
}

func initConfig() {
	if ConfigFilePath != "" {
		log.Debug("Config Path:", ConfigFilePath)
		if err := config.LoadConfig(ConfigFilePath); err != nil {
			log.Warn("Error while loading the configuration file. Loading the defaults")
		}
	}
}

func Execute() error {
	return generateCmd.Execute()
}
