package util

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
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

func Test_GetLatestTag(t *testing.T) {
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
					require.ErrorContains(t, err, tc.errMsg)
				} else {
					require.NoError(t, err)
				}

				if tc.isEmpty {
					require.Empty(t, tag)
				} else {
					require.NotEmpty(t, tag)
				}
			},
		)
	}
}
