package generator

import (
	"io"
	"strings"
	"text/template"
)

// Generator keeps the state of the generator
// where enabled plugin with the respective config is kept
type Generator struct {
	TrivyPluginEnabled bool
	TPConfig           TrivyPluginConfig
}

// TrivyPluginConfig stores the various configurations for trivy plugin
type TrivyPluginConfig struct {
	ExitCode          int    `mapstructure:"exit-code"`
	Timeout           string `mapstructure:"timeout"`
	Severity          string `mapstructure:"severity"`
	IgnoreUnfixed     bool   `mapstructure:"ignore-unfixed"`
	SecurityChecks    string `mapstructure:"security-checks"`
	SkipFiles         string `mapstructure:"skip-files"`
	SkipDirs          string `mapstructure:"skip-dirs"`
	ImageRef          string `mapstructure:"image-ref"`
	TrivyVersion      string `mapstructure:"version"`
	HelmOverridesFile string `mapstructure:"helm-overrides-file"`
}

// GenerateTrivyStep takes trivy plugin version and shell plugin version
// and an io.Writer to generate trivy step configuration. The trivy step is
// written to the provided io.Writer.
// It returns error in case write to the io.Writer errors out.
func GenerateTrivyStep(g Generator, w io.Writer, templateFolderPath string) error {
	funcMap := template.FuncMap{
		"join": func(arr []string) string {
			return strings.Join(arr, ",")
		},
	}

	tpl, err := template.New("").Funcs(funcMap).ParseGlob(templateFolderPath)
	if err != nil {
		return err
	}
	return tpl.ExecuteTemplate(w, "trivy-step.tmpl", g)
}
