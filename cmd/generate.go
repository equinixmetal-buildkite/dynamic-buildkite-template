package cmd

import (
	"dynamic-buildkite-template/generator"

	"github.com/spf13/cobra"
)

var (
	g = generator.Generator{}
)

func init() {
}

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

func Execute() error {
	return generateCmd.Execute()
}
