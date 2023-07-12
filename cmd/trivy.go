package cmd

import (
	"dynamic-buildkite-template/generator"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

var (
	trivyPluginConfig generator.TrivyPluginConfig
)

// LoadTrivyConfigs loads the trivy plugin config using "plugins.trivy" key
func LoadTrivyConfigs() {
	// load from config
	s := viper.Sub("plugins.trivy")
	if s == nil {
		log.Warn("Trivy Plugin configuration not found in the config file. .")
		return
	}

	err := s.Unmarshal(&trivyPluginConfig) // unmarshal to the trivyPluginConfig object
	if err != nil {
		log.Error("Error unmarshalling config file", err)
		return
	}

	// fetch latest trivy plugin version, if not defined in the config
	if strings.TrimSpace(trivyPluginConfig.TrivyVersion) == "" {
		trivyPluginConfig.TrivyVersion = GetLatestPluginTag("trivy-buildkite-plugin")
	}
	g.TPConfig = trivyPluginConfig
	// mark trivy plugin as enabled
	g.TrivyPluginEnabled = true
}
