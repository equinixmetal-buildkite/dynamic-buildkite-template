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
	generator.GenerateTrivyStep(g, &sb, "../../templates/*")

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
	err := generator.GenerateTrivyStep(g, &sb, "../templates/*") // wrong template path

	require.ErrorContains(t, err, "template: pattern matches no files", "Error generating trivy step")
}
