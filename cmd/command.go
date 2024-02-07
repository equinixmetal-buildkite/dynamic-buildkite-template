package cmd

import (
	"github.com/spf13/cobra"
)

// LoadCosignConfigs loads cosign plugin configuration from conf.yaml using "plugins.cosign" key
func LoadCommandConfigs(cmd *cobra.Command) {
	m, _ := cmd.Flags().GetStringToString("overrides")
	for k, v := range m {
		if k == "command" {
			g.CommandConfig.CommandConfig = v
			g.CommandConfigEnable = true
		}
	}
}
