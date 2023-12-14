package cmd

import (
	"dynamic-buildkite-template/generator"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

var (
	ssmPluginConfig generator.SSMPluginConfig
)

// LoadDockerMetaDataConfigs loads ssm-buildkite plugin configuration from conf.yaml using "plugins.dockermetadata" key
func LoadSSMDataConfigs() {
	// load from config
	s := viper.Sub("plugins.ssm-buildkite-plugin")
	if s == nil {
		log.Warn("ssm-buildkite Plugin configuration not found in the config file. .")
		return
	}

	log.Info("ssm-buildkite plugin found in the config file")

	err := s.Unmarshal(&ssmPluginConfig) // unmarshal to the dockermetadataPluginConfig object
	if err != nil {
		log.Error("Error unmarshalling ssm-buildkite plugin from config file", err)
		return
	}

	// fetch latest ssm-buildkite plugin version, if not defined in the config
	if strings.TrimSpace(dockermetadaPluginConfig.Version) == "" {
		dockermetadaPluginConfig.Version = GetLatestPluginTag("ssm-buildkite-plugin")
	}
	g.SSMConfig = ssmPluginConfig
	// mark ssm-buildkite plugin as enabled
	g.SSMPluginEnabled = true
}
