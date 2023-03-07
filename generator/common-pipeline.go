package generator

import (
	"fmt"
	"io"
)

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
	_, err := w.Write([]byte(trivyStep))
	return fmt.Errorf("error writing trivy step to the output stream: %W", err)
}
