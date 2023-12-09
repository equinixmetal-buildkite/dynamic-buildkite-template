package generator

import (
	"strings"
	"testing"
)

type testCase struct {
	name         string
	tpc          TrivyPluginConfig
	templatePath string
	hasError     bool
	errMsg       string
	expVal       string
}

func TestGenerateTrivyStep(t *testing.T) {
	testSeverity := "CRITICAL,HIGH"
	testSecurityChecks := "config,secret,vuln"
	testIgnoreUnfixed := true
	testSkipFiles := "cosign.key"

	tpc := TrivyPluginConfig{
		Severity:       &testSeverity,
		SecurityChecks: &testSecurityChecks,
		IgnoreUnfixed:  &testIgnoreUnfixed,
		SkipFiles:      &testSkipFiles,
	}
	expected := `
steps:
  - command: ls
    plugins:
      - equinixmetal-buildkite/trivy#<nil>:
          severity: "CRITICAL,HIGH"
          ignore-unfixed: true
          security-checks: "config,secret,vuln"
          skip-files: "cosign.key"
`
	cases := []testCase{
		{"success", tpc, "../templates/plugins-step.tmpl", false, "", expected},
		{"wrong_template_path", tpc, "../templates/xyz.tmpl", true, "no such file or directory", ""},
	}

	g := Generator{
		TrivyPluginEnabled: true,
		TPConfig:           tpc,
	}

	for _, tc := range cases {
		t.Run(
			tc.name,
			func(t *testing.T) {
				var sb strings.Builder
				err := GenerateBuildSteps(g, &sb, tc.templatePath)
				if tc.hasError {
					if !strings.Contains(err.Error(), tc.errMsg) {
						t.Fatalf("Error %s does not contain %s\n", err.Error(), tc.errMsg)
					}
				} else {
					if strings.TrimSpace(tc.expVal) != strings.TrimSpace(sb.String()) {
						t.Fatalf("Not equal: \n"+
							"expected: %s\n"+
							"actual  : %s\n", strings.TrimSpace(tc.expVal), strings.TrimSpace(sb.String()))
					}
				}
			},
		)
	}
}

type dockerTestCase struct {
	name         string
	dmpc         DockerMetadataPluginConfig
	templatePath string
	hasError     bool
	errMsg       string
	expVal       string
}

func TestGenerateDockerMetadataStep(t *testing.T) {
	images := []string{"test_image"}
	title := "test_title"
	vendor := "test_vendor"

	dmpc := DockerMetadataPluginConfig{
		Images: images,
		Title:  title,
		Vendor: vendor,
	}
	expected := `
steps:
  - command: ls
    plugins:
      - equinixmetal-buildkite/docker-metadata#:
          images:
          - "test_image"
          title: "test_title"
          vendor: "test_vendor"
`
	cases := []dockerTestCase{
		{"success", dmpc, "../templates/plugins-step.tmpl", false, "", expected},
	}

	g := Generator{
		DockerMetadataPluginEnabled: true,
		DockerMetadataConfig:        dmpc,
	}

	for _, tc := range cases {
		t.Run(
			tc.name,
			func(t *testing.T) {
				var sb strings.Builder
				err := GenerateBuildSteps(g, &sb, tc.templatePath)
				if tc.hasError {
					if !strings.Contains(err.Error(), tc.errMsg) {
						t.Fatalf("Error %s does not contain %s\n", err.Error(), tc.errMsg)
					}
				} else {
					if strings.TrimSpace(tc.expVal) != strings.TrimSpace(sb.String()) {
						t.Fatalf("Not equal: \n"+
							"expected: %s\n"+
							"actual  : %s\n", strings.TrimSpace(tc.expVal), strings.TrimSpace(sb.String()))
					}
				}
			},
		)
	}
}
