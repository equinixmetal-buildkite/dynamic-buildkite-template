package generator

// Generator keeps the state of the generator
// where enabled plugin with the respective config is kept
type Generator struct {
	TrivyPluginEnabled          bool
	CosignPluginEnabled         bool
	DockerMetadataPluginEnabled bool
	SSMPluginEnabled            bool
	DockerBuildPluginEnabled    bool
	CommandConfigEnable         bool
	TPConfig                    TrivyPluginConfig
	CosignConfig                CosignPluginConfig
	DockerMetadataConfig        DockerMetadataPluginConfig
	SSMConfig                   SSMPluginConfig
	DockerBuildConfig           DockerBuildConfig
	CommandConfig               CommandConfig
}
