package cmd

import (
	"dynamic-buildkite-template/generator"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

var (
	dockermetadaPluginConfig generator.DockerMetadataPluginConfig
)

// LoadDockerMetaDataConfigs loads docker metadata plugin configuration from conf.yaml using "plugins.dockermetadata" key
func LoadDockerMetaDataConfigs() {
	// load from config
	s := viper.Sub("plugins.docker-metadata")
	if s == nil {
		log.Warn("docker-metadata Plugin configuration not found in the config file. .")
		return
	}

	log.Info("docker-metadata plugin found in the config file")

	err := s.Unmarshal(&dockermetadaPluginConfig) // unmarshal to the dockermetadataPluginConfig object
	if err != nil {
		log.Error("Error unmarshalling docker-metadata plugin from config file", err)
		return
	}

	// fetch latest docker-metadata plugin version, if not defined in the config
	if strings.TrimSpace(dockermetadaPluginConfig.Version) == "" {
		dockermetadaPluginConfig.Version = GetLatestPluginTag("docker-metadata-buildkite-plugin")
	}
	g.DockerMetadataConfig = dockermetadaPluginConfig
	// mark docker-metadata plugin as enabled
	g.DockerMetadataPluginEnabled = true
}
