package generator

import (
	"io"
	"strings"
	"text/template"
)

// Generator is the struct to keep the values passed to the trivy step template
type Generator struct {
	TrivyPlugin    string
	ShellPlugin    string
	Severity       []string
	IgnoreUnfixed  bool
	SecurityChecks []string
	SkipFiles      string
}

// GenerateTrivyStep takes trivy plugin version and shell plugin version
// and an io.Writer to generate trivy step configuration. The trivy step is
// written to the provided io.Writer.
// It returns error in case write to the io.Writer errors out.
func GenerateTrivyStep(trivyPlugin, shellPlugin string, w io.Writer) error {
	generator := Generator{
		TrivyPlugin:    trivyPlugin,
		ShellPlugin:    shellPlugin,
		Severity:       []string{"CRITICAL", "HIGH"},
		IgnoreUnfixed:  true,
		SecurityChecks: []string{"config", "secret", "vuln"},
		SkipFiles:      "cosign.key",
	}

	funcMap := template.FuncMap{
		"join": func(arr []string) string {
			return strings.Join(arr, ",")
		},
	}

	tpl := template.Must(template.New("").Funcs(funcMap).ParseGlob("templates/*"))
	return tpl.ExecuteTemplate(w, "trivy-step.tmpl", generator)
}
