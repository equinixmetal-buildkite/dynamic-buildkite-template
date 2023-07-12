package cmd

import (
	"dynamic-buildkite-template/generator"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

var (
	cosignPluginConfig generator.CosignPluginConfig
)

// LoadCosignConfigs loads cosign plugin configuration from conf.yaml using "plugins.cosign" key
func LoadCosignConfigs() {
	// load from config
	s := viper.Sub("plugins.cosign")
	if s == nil {
		log.Warn("Cosign Plugin configuration not found in the config file. .")
		return
	}

	log.Info("Cosign plugin found in the config file")

	err := s.Unmarshal(&cosignPluginConfig) // unmarshal to the cosignPluginConfig object
	if err != nil {
		log.Error("Error unmarshalling config file", err)
		return
	}

	// fetch latest cosign plugin version, if not defined in the config
	if strings.TrimSpace(cosignPluginConfig.CosignVersion) == "" {
		cosignPluginConfig.CosignVersion = GetLatestPluginTag("cosign-buildkite-plugin")
	}
	g.CosignConfig = cosignPluginConfig
	// mark cosign plugin as enabled
	g.CosignPluginEnabled = true
}
