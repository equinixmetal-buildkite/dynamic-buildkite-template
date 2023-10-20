package cmd

import (
	"dynamic-buildkite-template/util"
	"os"

	log "github.com/sirupsen/logrus"
)

func GetLatestPluginTag(repoName string) string {
	githubPAT := os.Getenv("GITHUB_PAT")
	githubOrg := "equinixmetal-buildkite"
	tag, err := util.GetLatestTag(githubPAT, githubOrg, repoName)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("Latest %s plugin tag: %s", repoName, tag)
	return tag
}
