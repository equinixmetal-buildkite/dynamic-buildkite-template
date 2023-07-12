package generator_test

import (
	"dynamic-buildkite-template/generator"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GenerateSuccess(t *testing.T) {
	tpc := generator.TrivyPluginConfig{
		Severity:       "CRITICAL,HIGH",
		SecurityChecks: "config,secret,vuln",
		IgnoreUnfixed:  true,
		SkipFiles:      "cosign.key",
	}

	g := generator.Generator{
		TrivyPluginEnabled: true,
		TPConfig:           tpc,
	}

	var sb strings.Builder
	generator.GenerateBuildSteps(g, &sb, "../../templates/plugins-step.tmpl")

	expected := `
steps:
  - command: ls
    plugins:
      - equinixmetal-buildkite/trivy#:
          severity: CRITICAL,HIGH
          ignore-unfixed: true
          security-checks: config,secret,vuln
          skip-files: 'cosign.key'
`
	require.Equal(t, strings.TrimSpace(expected), strings.TrimSpace(sb.String()), "Generated template does not match")
}

func Test_GenerateWrongTemplatePath(t *testing.T) {
	tpc := generator.TrivyPluginConfig{
		Severity:       "CRITICAL,HIGH",
		SecurityChecks: "config,secret,vuln",
		IgnoreUnfixed:  true,
		SkipFiles:      "cosign.key",
	}

	g := generator.Generator{
		TrivyPluginEnabled: true,
		TPConfig:           tpc,
	}

	var sb strings.Builder
	err := generator.GenerateBuildSteps(g, &sb, "../templates/plugins-step.tmpl") // wrong template path

	require.ErrorContains(t, err, "The system cannot find the path specified.", "Error generating trivy step")
}
