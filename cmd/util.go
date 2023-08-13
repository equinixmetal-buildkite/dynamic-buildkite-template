package cmd

import (
	"dynamic-buildkite-template/util"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

func setFromStringFlag(f *string, cmd *cobra.Command, name string, doLookup bool) {
	// if doLookup is set then it would check for command line overrides before overriding the configuration
	// if doLookup is not set then it would pick the from default command line flag values
	if doLookup {
		if cmd.Flags().Lookup(name).Changed {
			*f = mustGetStringFlag(cmd, name)
		}
	} else {
		*f = mustGetStringFlag(cmd, name)
	}
}

func setFromBoolFlag(f *bool, cmd *cobra.Command, name string, doLookup bool) {
	if doLookup {
		if cmd.Flags().Lookup(name).Changed {
			*f = mustGetBoolFlag(cmd, name)
		}
	} else {
		*f = mustGetBoolFlag(cmd, name)
	}
}

func setFromIntFlag(f *int, cmd *cobra.Command, name string, doLookup bool) {
	if doLookup {
		if cmd.Flags().Lookup(name).Changed {
			*f = mustGetIntFlag(cmd, name)
		}
	} else {
		*f = mustGetIntFlag(cmd, name)
	}
}

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

// GetLatestPluginTag fetches the latest release tag of a plugin based on the repoName passed to it
func GetLatestPluginTag(repoName string) string {
	githubPAT := os.Getenv("GITHUB_PAT") // it looks for Github Personal Access Token
	githubOrg := "equinixmetal-buildkite"
	tag, err := util.GetLatestTag(githubPAT, githubOrg, repoName)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("Latest %s plugin tag: %s", repoName, tag)
	return tag
}
