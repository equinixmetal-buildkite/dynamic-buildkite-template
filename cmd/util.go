package cmd

import (
	"dynamic-buildkite-template/util"
	"log"
	"os"

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

func GetLatestTrivyPluginTag() string {
	githubPAT := os.Getenv("GITHUB_PAT")
	githubOrg := "equinixmetal-buildkite"
	repo := "trivy-buildkite-plugin"
	tag, err := util.GetLatestTag(githubPAT, githubOrg, repo)
	if err != nil {
		log.Fatal(err)
	}
	return tag
}
