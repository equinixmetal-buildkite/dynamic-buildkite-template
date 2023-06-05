package util

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/thedevsaddam/gojsonq/v2"
)

func GetLatestTag(githubPAT, githubOrg, repo string) (string, error) {
	githubAPIURL := fmt.Sprintf("https://%s:@api.github.com/repos/%s/%s", githubPAT, githubOrg, repo)

	latestURL := fmt.Sprintf("%s/releases/latest", githubAPIURL)
	client := resty.New()
	resp, err := client.R().Get(latestURL)
	if err != nil {
		return "", fmt.Errorf("Error while fetching latest tag: %w", err)
	}
	return gojsonq.New().FromString(resp.String()).Find("tag_name").(string), nil
}
