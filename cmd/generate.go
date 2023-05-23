package cmd

import (
	"dynamic-buildkite-template/generator"
	"os"

	"github.com/spf13/cobra"
)

var (
	g = generator.Generator{}
)

func init() {
	generateCmd.Flags().Bool("trivy-enabled", false, "Controls whether to generate step for trivy buildkite plugin")
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates plugin step for the provided plugins with configurations",
	Long: `
Usage of dynamic-buildkite-template
This Program generates step for the provided plugins with configurations
	`,
	// Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		CreateGenerator(cmd, args)
	},
}

func CreateGenerator(cmd *cobra.Command, args []string) {
	if mustGetBoolFlag(cmd, "trivy-enabled") {
		g.TrivyPluginEnabled = true
		trivyPluginConfig = generator.TrivyPluginConfig{
			ExitCode:          mustGetIntFlag(cmd, flagPrefix+"exit-code"),
			Timeout:           mustGetStringFlag(cmd, flagPrefix+"timeout"),
			Severity:          mustGetStringFlag(cmd, flagPrefix+"severity"),
			IgnoreUnfixed:     mustGetBoolFlag(cmd, flagPrefix+"ignore-unfixed"),
			SecurityChecks:    mustGetStringFlag(cmd, flagPrefix+"security-checks"),
			SkipFiles:         mustGetStringFlag(cmd, flagPrefix+"skip-files"),
			SkipDirs:          mustGetStringFlag(cmd, flagPrefix+"skip-dirs"),
			ImageRef:          mustGetStringFlag(cmd, flagPrefix+"image-ref"),
			TrivyVersion:      mustGetStringFlag(cmd, flagPrefix+"trivy-version"),
			HelmOverridesFile: mustGetStringFlag(cmd, flagPrefix+"helm-overrides-file"),
		}
		g.TPConfig = trivyPluginConfig
		generator.GenerateTrivyStep(g, os.Stdout, "templates/*")
	}
}

func Execute() error {
	return generateCmd.Execute()
}
