package generator

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
