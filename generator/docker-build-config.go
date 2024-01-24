package generator

// DockerBuildConfig stores the various configurations for docker build plugin
type DockerBuildConfig struct {
	Dockerfile string   `mapstructure:"dockerfile"`
	Context    string   `mapstructure:"context"`
	SecretFile string   `mapstructure:"secret-file"`
	Tags       []string `mapstructure:"tags"`
	Labels     []string `mapstructure:"labels"`
	BuildArgs  []string `mapstructure:"build-args"`
	Push       bool     `mapstructure:"push"`
	Version    string
}
