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

func mustGetStringArrayFlag(cmd *cobra.Command, name string) []string {
	flagVal, err := cmd.Flags().GetStringArray(name)
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
