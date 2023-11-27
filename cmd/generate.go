package cmd

import (
	"dynamic-buildkite-template/config"
	"dynamic-buildkite-template/generator"
	"dynamic-buildkite-template/util"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	g                     = generator.Generator{}
	ConfigFilePath        string
	defaultConfigFilePath = "conf.yaml"
)

func init() {
	cobra.OnInitialize(initConfig)
	// set the flag for passing overrides
	generateCmd.Flags().StringToString("overrides", nil, `pass the overrides in the maps syntax as --overrides plugins.trivy.skip-files="x.txt,y.txt" --overrides plugins.cosign.keyless=false`)

	generateCmd.PersistentFlags().StringVar(&ConfigFilePath, "config", "", fmt.Sprintf("config file path (default %q)", defaultConfigFilePath))
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates plugin step for the provided plugins with configurations",
	Long: `
Usage of dynamic-buildkite-template
This Program generates step for the provided plugins with configurations
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// check for overrides
		ParseOverrides(g, cmd)
		// load trivy plugin config
		LoadTrivyConfigs()
		// load cosign plugin config
		LoadCosignConfigs()
		// load docker-metadata plugin config
		LoadDockerMetaDataConfigs()
		// generate the build template
		err := generator.GenerateBuildSteps(g, os.Stdout, util.TemplateFilePath)
		if err != nil {
			log.Fatalf("Failed to generate build steps. %s", err.Error())
		}
	},
}

func initConfig() {
	if ConfigFilePath != "" {
		log.Debug("config path: ", ConfigFilePath)
		if err := config.LoadConfig(ConfigFilePath); err != nil {
			log.Fatal("error while loading the configuration file. Exiting the program.")
		}
		return
	}

	log.Debug("config path:", defaultConfigFilePath)
	if err := config.LoadConfig(defaultConfigFilePath); err != nil {
		log.Fatalf("error while loading the configuration file: %s. Configuration file must be present.", defaultConfigFilePath)
	}
}

func Execute() error {
	return generateCmd.Execute()
}

// ParseOverrides checks for command line flags for the overrides and updates the viper global object
func ParseOverrides(g generator.Generator, cmd *cobra.Command) {
	m, err := cmd.Flags().GetStringToString("overrides") // check for --overrides flag for map[string]string
	if err != nil {
		log.Warn("No overrides defined. Continuing with defaults defined in the config file.")
		return
	}
	vNew := viper.New() // new viper object for storing overrides
	for k, v := range m {
		vNew.Set(k, v)
	}

	err = viper.MergeConfigMap(vNew.AllSettings()) // merge to global viper object
	if err != nil {
		log.Fatalf("Failed merging the configuration with command line overrides. %s", err.Error())
	}
}
