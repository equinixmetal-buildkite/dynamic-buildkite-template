package cmd

import (
	"dynamic-buildkite-template/generator"
)

var (
	trivyPluginConfig generator.TrivyPluginConfig
	flagPrefix        = "trivy-"
)

func init() {
	generateCmd.Flags().Int(flagPrefix+"exit-code", 0, "Controls whether the security scan is blocking or not for trivy buildkite plugin")
	generateCmd.Flags().String(flagPrefix+"timeout", "", "Controls the maximum amount of time a scan will run for trivy buildkite plugin")
	generateCmd.Flags().String(flagPrefix+"severity", "UNKNOWN,LOW,MEDIUM,HIGH,CRITICAL", "Controls the severity of the vulnerabilities to be scanned for trivy buildkite plugin")
	generateCmd.Flags().Bool(flagPrefix+"ignore-unfixed", false, "Controls whether to display only fixed vulnerabilities for trivy buildkite plugin")
	generateCmd.Flags().String(flagPrefix+"security-checks", "vuln,config", "Controls the security checks to be performed for trivy buildkite plugin")
	generateCmd.Flags().String(flagPrefix+"skip-files", "", "Controls the files to be skipped during the scan for trivy buildkite plugin")
	generateCmd.Flags().String(flagPrefix+"skip-dirs", "", "Controls the directories to be skipped during the scan for trivy buildkite plugin")
	generateCmd.Flags().String(flagPrefix+"image-ref", "", "Controls the image reference to be scanned for trivy buildkite plugin")
	generateCmd.Flags().String(flagPrefix+"trivy-version", "", "Controls the version of trivy to be used for trivy buildkite plugin")
	generateCmd.Flags().String(flagPrefix+"helm-overrides-file", "", "To pass helm override values to trivy config scan for trivy buildkite pluginn")
}
