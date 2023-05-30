package cmd

import (
	"dynamic-buildkite-template/generator"
	"os"

	"github.com/spf13/cobra"
)

var (
	trivyPluginConfig generator.TrivyPluginConfig
)

var trivyCmd = &cobra.Command{
	Use:   "trivy",
	Short: "Generates trivy plugin step for the given configurations",
	Long: `
Generates trivy plugin step for the given configurations
	`,
	Run: func(cmd *cobra.Command, args []string) {
		CreateGenerator(cmd, args)
	},
}

func init() {

	trivyCmd.Flags().Int("exit-code", 0, "Controls whether the security scan is blocking or not for trivy buildkite plugin")
	trivyCmd.Flags().String("timeout", "15m", "Controls the maximum amount of time a scan will run for trivy buildkite plugin")
	trivyCmd.Flags().String("severity", "HIGH,CRITICAL", "Controls the severity of the vulnerabilities to be scanned for trivy buildkite plugin")
	trivyCmd.Flags().Bool("ignore-unfixed", true, "Controls whether to display only fixed vulnerabilities for trivy buildkite plugin")
	trivyCmd.Flags().String("security-checks", "vuln,config", "Controls the security checks to be performed for trivy buildkite plugin")
	trivyCmd.Flags().String("skip-files", "", "Controls the files to be skipped during the scan for trivy buildkite plugin")
	trivyCmd.Flags().String("skip-dirs", "", "Controls the directories to be skipped during the scan for trivy buildkite plugin")
	trivyCmd.Flags().String("image-ref", "", "Controls the image reference to be scanned for trivy buildkite plugin")
	trivyCmd.Flags().String("version", "v1.18.2", "Controls the version of trivy to be used for trivy buildkite plugin")
	trivyCmd.Flags().String("helm-overrides-file", "", "To pass helm override values to trivy config scan for trivy buildkite pluginn")

	generateCmd.AddCommand(trivyCmd)
}

// CreateGenerator populates the TrivyPluginConfig by reading the values from the command line flags
func CreateGenerator(cmd *cobra.Command, args []string) {
	g.TrivyPluginEnabled = true
	trivyPluginConfig = generator.TrivyPluginConfig{
		ExitCode:          mustGetIntFlag(cmd, "exit-code"),
		Timeout:           mustGetStringFlag(cmd, "timeout"),
		Severity:          mustGetStringFlag(cmd, "severity"),
		IgnoreUnfixed:     mustGetBoolFlag(cmd, "ignore-unfixed"),
		SecurityChecks:    mustGetStringFlag(cmd, "security-checks"),
		SkipFiles:         mustGetStringFlag(cmd, "skip-files"),
		SkipDirs:          mustGetStringFlag(cmd, "skip-dirs"),
		ImageRef:          mustGetStringFlag(cmd, "image-ref"),
		TrivyVersion:      mustGetStringFlag(cmd, "version"),
		HelmOverridesFile: mustGetStringFlag(cmd, "helm-overrides-file"),
	}
	g.TPConfig = trivyPluginConfig
	generator.GenerateTrivyStep(g, os.Stdout, "templates/*")
}
