package util

import (
	"os"
	"strings"
	"testing"
)

type testCase struct {
	name     string
	input    gitDetails
	hasError bool
	errMsg   string
	isEmpty  bool
}

type gitDetails struct {
	PAT  string
	Org  string
	Repo string
}

func TestGetLatestTag(t *testing.T) {
	cases := []testCase{
		{
			"PAT_provided",
			gitDetails{os.Getenv("GITHUB_PAT"), "equinixmetal-buildkite", "trivy-buildkite-plugin"},
			false,
			"",
			false,
		},
		{
			"PAT_absent",
			gitDetails{"", "equinixmetal-buildkite", "trivy-buildkite-plugin"},
			true,
			"error while fetching latest tag",
			true,
		},
	}

	for _, tc := range cases {
		t.Run(
			tc.name,
			func(t *testing.T) {
				tag, err := GetLatestTag(tc.input.PAT, tc.input.Org, tc.input.Repo)

				if tc.hasError {
					if !strings.Contains(err.Error(), tc.errMsg) {
						t.Fatalf("Error %s does not contain %s\n", err.Error(), tc.errMsg)
					}
				} else {
					if err != nil {
						t.Fatalf("Error %s is not nil", err.Error())
					}
				}

				if tc.isEmpty {
					if strings.TrimSpace(tag) != "" {
						t.Fatalf("Tag expected to be blank but has value: %s", tag)
					}
				} else {
					if strings.TrimSpace(tag) == "" {
						t.Fatal("Tag expected to have value but is blank")
					}
				}
			},
		)
	}
}
