package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// GithubLatestReleaseResponse is used to unmarshal the response while fetching the github latest release
type GithubLatestReleaseResponse struct {
	TagName string `json:"tag_name"`
}

// GetLatestTag fetches the latest released tag from a github repo
func GetLatestTag(gitToken, githubOrg, repo string) (string, error) {
	// we can move timeout to conf.yaml as configuration
	// as part of https://github.com/equinixmetal-buildkite/dynamic-buildkite-template/pull/19/files
	timeout := 15 * time.Second

	githubAPIURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", githubOrg, repo)

	req, err := http.NewRequest(http.MethodGet, githubAPIURL, nil)
	if err != nil {
		return "", fmt.Errorf("error while creating request to fetch the latest tag: %w", err)
	}
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", gitToken))
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")

	client := http.Client{
		Timeout: timeout,
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error while sending request to fetch the latest tag: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error while fetching latest tag. Status Code: %d", resp.StatusCode)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body while fetching latest tag: %w", err)
	}

	var ghLastestReleaseResp GithubLatestReleaseResponse
	err = json.Unmarshal(b, &ghLastestReleaseResp)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling response while fetching latest tag: %w", err)
	}

	return ghLastestReleaseResp.TagName, nil
}
