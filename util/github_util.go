package util

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"

	"github.com/thedevsaddam/gojsonq/v2"
)

// GetLatestTag fetches the latest release tag using Github REST API
func GetLatestTag(githubPAT, githubOrg, repo string) (string, error) {
	githubAPIURL := fmt.Sprintf("https://%s:@api.github.com/repos/%s/%s", githubPAT, githubOrg, repo)

	latestURL := fmt.Sprintf("%s/releases/latest", githubAPIURL)
	log.Debug("Latest tag url:", latestURL)
	client := resty.New()
	resp, err := client.R().Get(latestURL)
	if err != nil {
		return "", fmt.Errorf("error while fetching latest tag: %w", err)
	}
	return gojsonq.New().FromString(resp.String()).Find("tag_name").(string), nil
}
