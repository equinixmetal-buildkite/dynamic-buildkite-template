package generator

import (
	"fmt"
	"io"
)

// GenerateTrivyStep takes trivy plugin version and shell plugin version
// and an io.Writer to generate trivy step configuration. The trivy step is
// written to the provided io.Writer.
// It returns error in case write to the io.Writer errors out.
func GenerateTrivyStep(trivyPlugin, shellPlugin string, w io.Writer) error {
	trivyStepFormat := `
steps:
- command: ls
  plugins:
    - equinixmetal-buildkite/trivy#%s
- label: ":sparkles: SHELL CHECK"
  plugins:
    - shellcheck#%s:
        files: script.sh
`

	trivyStep := fmt.Sprintf(trivyStepFormat, trivyPlugin, shellPlugin)
	if _, err := w.Write([]byte(trivyStep)); err != nil {
		return fmt.Errorf("error writing trivy step to the output stream: %W", err)
	}
	return nil
}
