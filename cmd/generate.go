package cmd

import (
	"dynamic-buildkite-template/generator"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	generateCmd.PersistentFlags().String("trivyPlugin", "", "provide trivy plugin version")
	generateCmd.PersistentFlags().String("shellPlugin", "", "provide shell plugin version")
	generateCmd.PersistentFlags().Bool("ignoreUnfixed", true, "provide if unfixed items are to be ignored")
	generateCmd.PersistentFlags().String("skipFiles", "cosign.key", "provide files to be skipped in trivy plugin")
	generateCmd.PersistentFlags().String("shellCheckFiles", "script.sh", "provide files to be checked by the shell plugin")

	generateCmd.PersistentFlags().StringArray("severity", []string{"CRITICAL", "HIGH"}, "provide the severity")
	generateCmd.PersistentFlags().StringArray("securityChecks", []string{"config", "secret", "vuln"}, "provide the security checks")
}

func Execute() error {
	return generateCmd.Execute()
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates trivy step for the provided options",
	Long: `
Usage of dynamic-buildkite-template
This Program generates trivy step for the provided options
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var g = generator.Generator{
			TrivyPlugin:     mustGetStringFlag(cmd, "trivyPlugin"),
			ShellPlugin:     mustGetStringFlag(cmd, "shellPlugin"),
			Severity:        mustGetStringArrayFlag(cmd, "severity"),
			IgnoreUnfixed:   mustGetBoolFlag(cmd, "ignoreUnfixed"),
			SecurityChecks:  mustGetStringArrayFlag(cmd, "securityChecks"),
			SkipFiles:       mustGetStringFlag(cmd, "skipFiles"),
			ShellCheckFiles: mustGetStringFlag(cmd, "shellCheckFiles"),
		}
		generator.GenerateTrivyStep(g, os.Stdout, "templates/*")
	},
}
