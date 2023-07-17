package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// GithubLatestReleaseResponse is used to unmarshal the response while fetching the github latest release
type GithubLatestReleaseResponse struct {
	URL       string `json:"url"`
	AssetsURL string `json:"assets_url"`
	UploadURL string `json:"upload_url"`
	HTMLURL   string `json:"html_url"`
	ID        int    `json:"id"`
	Author    struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"author"`
	NodeID          string    `json:"node_id"`
	TagName         string    `json:"tag_name"`
	TargetCommitish string    `json:"target_commitish"`
	Name            string    `json:"name"`
	Draft           bool      `json:"draft"`
	Prerelease      bool      `json:"prerelease"`
	CreatedAt       time.Time `json:"created_at"`
	PublishedAt     time.Time `json:"published_at"`
	Assets          []any     `json:"assets"`
	TarballURL      string    `json:"tarball_url"`
	ZipballURL      string    `json:"zipball_url"`
	Body            string    `json:"body"`
	MentionsCount   int       `json:"mentions_count"`
}

// GetLatestTag fetches the latest released tag from a github repo
func GetLatestTag(githubPAT, githubOrg, repo string) (string, error) {
	githubAPIURL := fmt.Sprintf("https://%s:@api.github.com/repos/%s/%s", githubPAT, githubOrg, repo)

	latestURL := fmt.Sprintf("%s/releases/latest", githubAPIURL)

	resp, err := http.Get(latestURL)
	if err != nil {
		return "", fmt.Errorf("error while fetching latest tag: %w", err)
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
	decoder := json.NewDecoder(strings.NewReader(string(b)))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&ghLastestReleaseResp)
	// err = json.Unmarshal(b, &ghLastestReleaseResp)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling response while fetching latest tag: %w", err)
	}

	return ghLastestReleaseResp.TagName, nil
}
