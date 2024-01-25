package cmd

import (
	"dynamic-buildkite-template/generator"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

var (
	dockerBuildPluginConfig generator.DockerBuildConfig
)

// LoadDockerBuildConfigs loads docker build plugin configuration from conf.yaml using "plugins.docker-build" key
func LoadDockerBuildConfigs() {
	// load from config
	s := viper.Sub("plugins.docker-build")
	if s == nil {
		log.Warn("Docker Build Plugin configuration not found in the config file. .")
		return
	}

	log.Info("Docker Build plugin found in the config file")

	err := s.Unmarshal(&dockerBuildPluginConfig) // unmarshal to the dockerBuildPluginConfig object
	if err != nil {
		log.Error("Error unmarshalling docker plugin from config file", err)
		return
	}

	// fetch latest docker build plugin version, if not defined in the config
	if strings.TrimSpace(dockerBuildPluginConfig.Version) == "" {
		dockerBuildPluginConfig.Version = GetLatestPluginTag("docker-build-buildkite-plugin")
	}
	g.DockerBuildConfig = dockerBuildPluginConfig
	// mark docker build plugin as enabled
	g.DockerBuildPluginEnabled = true
}
