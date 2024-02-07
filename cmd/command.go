package cmd

import (
	"github.com/spf13/cobra"
)

// var (
// 	commandConfig generator.CommandConfig
// )

// LoadCosignConfigs loads cosign plugin configuration from conf.yaml using "plugins.cosign" key
func LoadCommandConfigs(cmd *cobra.Command) {

	//g.CommandConfig.CommandConfig = "ls"
	m, _ := cmd.Flags().GetStringToString("overrides")
	for k, v := range m {
		if k == "command" {
			//fmt.Println("command key found!!!!", v)
			g.CommandConfig.CommandConfig = v
			g.CommandConfigEnable = true
		}
	}
}
