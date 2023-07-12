package generator

// CosignPluginConfig stores the various configurations for cosign plugin
type CosignPluginConfig struct {
	Image         string        `mapstructure:"image"`
	Keyless       bool          `mapstructure:"keyless"`
	KeylessConfig KeylessConfig `mapstructure:"keyless-config"`
	KeyedConfig   KeyedConfig   `mapstructure:"keyed-config"`
	CosignVersion string        `mapstructure:"cosign-version"`
}

// KeylessConfig is used if Keyless is set to true in CosignPluginConfig
type KeylessConfig struct {
	FulcioURL string `mapstructure:"fulcio_url"`
	RekorURL  string `mapstructure:"rekor_url"`
}

// KeyedConfig is used if Keyless is set to false in CosignPluginConfig
type KeyedConfig struct {
	Key string `mapstructure:"key"`
}
