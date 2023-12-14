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

type ssmTestCase struct {
	name         string
	ssm          SSMPluginConfig
	templatePath string
	hasError     bool
	errMsg       string
	expVal       string
}

func TestSSMStep(t *testing.T) {

	CosignKeySecret := "test-secret"
	CosignPasswd := "passwd"

	ssm := SSMPluginConfig{
		Parameter: Parameter{CosignKeySecret: CosignKeySecret, CosignPassword: CosignPasswd},
	}
	expected := `
steps:
  - command: ls
    plugins:
      - ssh://git@git@github.com/equinixmetal/ssm-buildkite-plugin#:
          parameters:
            COSIGN_KEY_SECRET : test-secret
            COSIGN_PASSWORD : passwd
`
	cases := []ssmTestCase{
		{"success", ssm, "../templates/plugins-step.tmpl", false, "", expected},
	}

	g := Generator{
		SSMPluginEnabled: true,
		SSMConfig:        ssm,
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

type cosignTestCase struct {
	name         string
	cosignConf   CosignPluginConfig
	templatePath string
	hasError     bool
	errMsg       string
	expVal       string
}

func TestCosignStep(t *testing.T) {

	image := "ghcr.io/my-project/my-image:latest"
	keyless := false
	key := "sample-key"

	cosignConfig := CosignPluginConfig{
		Image:       image,
		Keyless:     keyless,
		KeyedConfig: KeyedConfig{Key: key},
	}
	expected := `
steps:
  - label: ":docker: get cosign key"
    key: "getkey"
    command: |
      #!/bin/bash
      echo "\$COSIGN_KEY_SECRET" > ${COSIGN_KEY_PATH}
    plugins:
      - equinixmetal-buildkite/cosign#:
          image: ghcr.io/my-project/my-image:latest
          keyless : false
          keyed-config:
            key: sample-key
`
	cases := []cosignTestCase{
		{"success", cosignConfig, "../templates/plugins-step.tmpl", false, "", expected},
	}

	g := Generator{
		CosignPluginEnabled: true,
		CosignConfig:        cosignConfig,
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
