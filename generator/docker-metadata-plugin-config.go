package generator

type DockerMetadataPluginConfig struct {
	Images    []string `mapstructure:"images"`
	ExtraTags []string `mapstructure:"extra_tags"`
	Title     string   `mapstructure:"title"`
	Licenses  string   `mapstructure:"licenses"`
	Vendor    string   `mapstructure:"vendor"`
	Debug     bool     `mapstructure:"debug"`
	Version   string   `mapstructure:"docker-metadata-version"`
}
