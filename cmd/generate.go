package cmd

import (
	"dynamic-buildkite-template/generator"
	"dynamic-buildkite-template/util"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	g = generator.Generator{} // generator object containing all plugin configuration
)

func init() {
	// set the flag for passing overrides
	generateCmd.Flags().StringToString("overrides", nil, `pass the overrides in the maps syntax as --overrides plugins.trivy.skip-files="x.txt,y.txt" --overrides plugins.cosign.keyless=false`)
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
		// generate the build template
		generator.GenerateBuildSteps(g, os.Stdout, util.TEMPLATE_FILE_PATH)
	},
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

	viper.MergeConfigMap(vNew.AllSettings()) // merge to global viper object
}

func Execute() error {
	return generateCmd.Execute()
}
