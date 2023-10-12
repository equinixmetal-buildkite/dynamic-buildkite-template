package generator

import (
	"io"
	"os"
	"text/template"
)

// GenerateTrivyStep takes trivy plugin version and shell plugin version
// and an io.Writer to generate trivy step configuration. The trivy step is
// written to the provided io.Writer.
// It returns error in case write to the io.Writer errors out.
func GenerateBuildSteps(g Generator, w io.Writer, templateFilePath string) error {
	file, err := os.ReadFile(templateFilePath)
	if err != nil {
		return err
	}

	// Parse the template contents
	tmpl, err := template.New("tmpl").Parse(string(file))
	if err != nil {
		return err
	}

	// Use the template.Execute() function to apply the data object to the parsed template
	return tmpl.Execute(w, g)
}
