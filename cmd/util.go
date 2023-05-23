package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func mustGetStringFlag(cmd *cobra.Command, name string) string {
	flagVal, err := cmd.Flags().GetString(name)
	if err != nil {
		log.Fatalf("Failed to get value of %s. %s", name, err.Error())
	}
	return flagVal
}

func mustGetBoolFlag(cmd *cobra.Command, name string) bool {
	flagVal, err := cmd.Flags().GetBool(name)
	if err != nil {
		log.Fatalf("Failed to get value of %s. %s", name, err.Error())
	}
	return flagVal
}

func mustGetIntFlag(cmd *cobra.Command, name string) int {
	flagVal, err := cmd.Flags().GetInt(name)
	if err != nil {
		log.Fatalf("Failed to get value of %s. %s", name, err.Error())
	}
	return flagVal
}

func subCommandExists(cmd *cobra.Command, name string) bool {
	for _, cmd := range cmd.Commands() {
		if cmd.Name() == name {
			return true
		}
	}
	return false
}
