package util

import (
	"errors"
	"os"
	"testing"
)

func TestGetLatestTag(t *testing.T) {
	type repo struct {
		owner  string
		name   string
		public bool
	}

	repos := map[string]repo{
		"PublicRepo": {
			"equinixmetal-buildkite",
			"trivy-buildkite-plugin",
			true,
		},
		// because this is open-source, here we actually use same public repo to with test github pat provided via GITHUB_PAT env var
		"PrivateRepo": {
			"equinixmetal-buildkite",
			"trivy-buildkite-plugin",
			false,
		},
	}

	for n, r := range repos {
		t.Run(n, func(t *testing.T) {
			cases := []struct {
				name    string
				pat     string
				wantErr error
			}{
				{
					"ValidPAT",
					os.Getenv("GITHUB_PAT"),
					nil,
				},
				{
					"InvalidPAT",
					"this is not a valid gihub pat",
					ErrFetchingLatestTag,
				},
				{
					"NoPAT",
					"",
					func() error {
						if r.public {
							return nil
						}
						return ErrFetchingLatestTag
					}(),
				},
			}

			for _, tc := range cases {
				t.Run(tc.name, func(t *testing.T) {
					tag, err := GetLatestTag(tc.pat, r.owner, r.name)
					if err != nil {
						if tc.wantErr == nil {
							t.Fatalf("unexpected error: %v", err)
						}
						if !errors.Is(err, ErrFetchingLatestTag) {
							t.Fatalf("want: %v, got: %v", tc.wantErr, err)
						}
						return
					}
					if tag == "" {
						t.Fatalf("unexpected empty tag: %v", err)
					}
				},
				)
			}
		})
	}
}
