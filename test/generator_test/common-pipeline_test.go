package generator_test

import (
	"dynamic-buildkite-template/generator"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GenerateSuccess(t *testing.T) {
	g := generator.Generator{
		TrivyPlugin:     "v1.18.2",
		ShellPlugin:     "v1.2.5",
		Severity:        []string{"CRITICAL", "HIGH"},
		SecurityChecks:  []string{"config", "secret", "vuln"},
		IgnoreUnfixed:   true,
		SkipFiles:       "cosign.key",
		ShellCheckFiles: "script.sh",
	}

	var sb strings.Builder
	generator.GenerateTrivyStep(g, &sb, "../../templates/*")

	expected := "steps:\n- command: ls\n  plugins:\n    - equinixmetal-buildkite/trivy#v1.18.2\n\t    severity: CRITICAL,HIGH\n\t    ignore-unfixed: true\n\t    security-checks: config,secret,vuln\n\t    skip-files: 'cosign.key'\n\n- label: \":sparkles: SHELL CHECK\"\n  plugins:\n    - shellcheck#v1.2.5:\n        files: script.sh"
	require.Equal(t, strings.TrimSpace(expected), strings.TrimSpace(sb.String()), "Generated template does not match")
}

func Test_GenerateWrongTemplatePath(t *testing.T) {
	g := generator.Generator{
		TrivyPlugin:     "v1.18.2",
		ShellPlugin:     "v1.2.5",
		Severity:        []string{"CRITICAL", "HIGH"},
		SecurityChecks:  []string{"config", "secret", "vuln"},
		IgnoreUnfixed:   true,
		SkipFiles:       "cosign.key",
		ShellCheckFiles: "script.sh",
	}

	var sb strings.Builder
	err := generator.GenerateTrivyStep(g, &sb, "../templates/*") // wrong template path

	require.ErrorContains(t, err, "template: pattern matches no files", "Error generating trivy step")
}
