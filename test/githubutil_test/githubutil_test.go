package githubutil_test

import (
	"dynamic-buildkite-template/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GetLatestTagSuccess(t *testing.T) {
	githubPAT := "" // publis repos do not need PAT
	githubOrg := "equinixmetal-buildkite"
	repo := "trivy-buildkite-plugin"
	tag, err := util.GetLatestTag(githubPAT, githubOrg, repo)
	require.NoError(t, err, "Error fetching latest tag")

	require.NotEmpty(t, tag, "Blank tag received")
}

func Test_GetLatestTagError(t *testing.T) {
	githubPAT := "" // private repos need PAT
	githubOrg := "hnadimint"
	repo := "buildkite-actions"
	tag, err := util.GetLatestTag(githubPAT, githubOrg, repo)

	require.ErrorContains(t, err, "error while fetching latest tag")
	require.Empty(t, tag)
}
