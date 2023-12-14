package generator

type SSMPluginConfig struct {
	Parameter Parameter `mapstructure:"parameters"`
}

type Parameter struct {
	CosignKeySecret string `mapstructure:"COSIGN_KEY_SECRET"`
	CosignPassword  string `mapstructure:"COSIGN_PASSWORD"`
	GithubToken     string `mapstructure:"GITHUB_TOKEN"`
}
