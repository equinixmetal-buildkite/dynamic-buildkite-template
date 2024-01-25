package generator

import (
	"io"
	"os"
	"text/template"
)

// GenerateBuildSteps takes a Generator object, an io.Writer, and a templateFilePath
// to generate build step configuration. The build step is written to the provided io.Writer.
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
