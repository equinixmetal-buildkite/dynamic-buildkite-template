package cmd

import (
	"dynamic-buildkite-template/generator"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	// Defaults are placeholders here. Actual defaults are defined in the config file
	trivyCmd.Flags().Int("exit-code", 0, "Controls whether the security scan is blocking or not for trivy buildkite plugin")
	trivyCmd.Flags().String("timeout", "5m0s", "Controls the maximum amount of time a scan will run for trivy buildkite plugin")
	trivyCmd.Flags().String("severity", "HIGH,CRITICAL", "Controls the severity of the vulnerabilities to be scanned for trivy buildkite plugin")
	trivyCmd.Flags().Bool("ignore-unfixed", true, "Controls whether to display only fixed vulnerabilities for trivy buildkite plugin")
	trivyCmd.Flags().String("security-checks", "vuln,config", "Controls the security checks to be performed for trivy buildkite plugin")
	trivyCmd.Flags().String("skip-files", "", "Controls the files to be skipped during the scan for trivy buildkite plugin")
	trivyCmd.Flags().String("skip-dirs", "", "Controls the directories to be skipped during the scan for trivy buildkite plugin")
	trivyCmd.Flags().String("image-ref", "", "Controls the image reference to be scanned for trivy buildkite plugin")
	trivyCmd.Flags().String("version", "", "Controls the version of trivy to be used for trivy buildkite plugin")
	trivyCmd.Flags().String("helm-overrides-file", "", "To pass helm override values to trivy config scan for trivy buildkite pluginn")

	generateCmd.AddCommand(trivyCmd)
}

// CreateGenerator populates the TrivyPluginConfig by reading the values from the command line flags
func CreateGenerator(cmd *cobra.Command, args []string) {
	g.TrivyPluginEnabled = true
	// load from config
	s := viper.Sub("plugins.trivy")
	doLookup := true
	if s == nil {
		log.Debug("trivy plugin configuration not found in the config file or wrong config file. Proceeding with defaults from command line flags.")
		doLookup = false
	} else {
		err := s.Unmarshal(&trivyPluginConfig)
		if err != nil {
			log.Error("error unmarshalling config file", err)
		}
	}

	// initializing nil fields for the cases when conf file is not present or there is command line flag override
	if (!doLookup || cmd.Flags().Lookup("exit-code").Changed) && trivyPluginConfig.ExitCode == nil {
		trivyPluginConfig.ExitCode = new(int)
	}
	if (!doLookup || cmd.Flags().Lookup("timeout").Changed) && trivyPluginConfig.Timeout == nil {
		trivyPluginConfig.Timeout = new(string)
	}
	if (!doLookup || cmd.Flags().Lookup("severity").Changed) && trivyPluginConfig.Severity == nil {
		trivyPluginConfig.Severity = new(string)
	}
	if (!doLookup || cmd.Flags().Lookup("ignore-unfixed").Changed) && trivyPluginConfig.IgnoreUnfixed == nil {
		trivyPluginConfig.IgnoreUnfixed = new(bool)
	}
	if (!doLookup || cmd.Flags().Lookup("security-checks").Changed) && trivyPluginConfig.SecurityChecks == nil {
		trivyPluginConfig.SecurityChecks = new(string)
	}
	if (!doLookup || cmd.Flags().Lookup("skip-files").Changed) && trivyPluginConfig.SkipFiles == nil {
		trivyPluginConfig.SkipFiles = new(string)
	}
	if (!doLookup || cmd.Flags().Lookup("skip-dirs").Changed) && trivyPluginConfig.SkipDirs == nil {
		trivyPluginConfig.SkipDirs = new(string)
	}
	if (!doLookup || cmd.Flags().Lookup("image-ref").Changed) && trivyPluginConfig.ImageRef == nil {
		trivyPluginConfig.ImageRef = new(string)
	}
	if (!doLookup || cmd.Flags().Lookup("version").Changed) && trivyPluginConfig.TrivyVersion == nil {
		trivyPluginConfig.TrivyVersion = new(string)
	}
	if (!doLookup || cmd.Flags().Lookup("helm-overrides-file").Changed) && trivyPluginConfig.HelmOverridesFile == nil {
		trivyPluginConfig.HelmOverridesFile = new(string)
	}

	setFromIntFlag(trivyPluginConfig.ExitCode, cmd, "exit-code", doLookup)
	setFromStringFlag(trivyPluginConfig.Timeout, cmd, "timeout", doLookup)
	setFromStringFlag(trivyPluginConfig.Severity, cmd, "severity", doLookup)
	setFromBoolFlag(trivyPluginConfig.IgnoreUnfixed, cmd, "ignore-unfixed", doLookup)
	setFromStringFlag(trivyPluginConfig.SecurityChecks, cmd, "security-checks", doLookup)
	setFromStringFlag(trivyPluginConfig.SkipFiles, cmd, "skip-files", doLookup)
	setFromStringFlag(trivyPluginConfig.SkipDirs, cmd, "skip-dirs", doLookup)
	setFromStringFlag(trivyPluginConfig.ImageRef, cmd, "image-ref", doLookup)
	setFromStringFlag(trivyPluginConfig.TrivyVersion, cmd, "version", doLookup)
	setFromStringFlag(trivyPluginConfig.HelmOverridesFile, cmd, "helm-overrides-file", doLookup)

	if trivyPluginConfig.TrivyVersion == nil {
		trivyPluginConfig.TrivyVersion = new(string)
	}
	if strings.TrimSpace(*trivyPluginConfig.TrivyVersion) == "" {
		lv := GetLatestTrivyPluginTag()
		trivyPluginConfig.TrivyVersion = &lv
	}
	g.TPConfig = trivyPluginConfig
	generator.GenerateTrivyStep(g, os.Stdout, "templates/*")
}
