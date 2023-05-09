package generator

import (
	"io"
	"strings"
	"text/template"
)

// Generator is the struct to keep the values passed to the trivy step template
type Generator struct {
	TrivyPlugin     string
	ShellPlugin     string
	Severity        []string
	IgnoreUnfixed   bool
	SecurityChecks  []string
	SkipFiles       string
	ShellCheckFiles string
}

// GenerateTrivyStep takes trivy plugin version and shell plugin version
// and an io.Writer to generate trivy step configuration. The trivy step is
// written to the provided io.Writer.
// It returns error in case write to the io.Writer errors out.
func GenerateTrivyStep(trivyPlugin, shellPlugin string, severity []string, ignoreUnfixed bool, securityChecks []string, skipFiles, shellCheckFiles string, w io.Writer) error {
	generator := Generator{}
	generator.TrivyPlugin = trivyPlugin
	generator.ShellPlugin = shellPlugin
	// initialize with defaults for arrays as arrays are using ArrayFlags type and the default cannot be initialized in flag.Var()
	generator.Severity = []string{"CRITICAL", "HIGH"}
	generator.SecurityChecks = []string{"config", "secret", "vuln"}

	// check for overrides
	if severity != nil {
		generator.Severity = severity
	}

	generator.IgnoreUnfixed = ignoreUnfixed

	if securityChecks != nil {
		generator.SecurityChecks = securityChecks
	}
	generator.SkipFiles = skipFiles
	generator.ShellCheckFiles = shellCheckFiles

	funcMap := template.FuncMap{
		"join": func(arr []string) string {
			return strings.Join(arr, ",")
		},
	}

	tpl := template.Must(template.New("").Funcs(funcMap).ParseGlob("templates/*"))
	return tpl.ExecuteTemplate(w, "trivy-step.tmpl", generator)
}
