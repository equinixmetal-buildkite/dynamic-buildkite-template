package config

import "github.com/spf13/viper"

func LoadConfig(configFilePath string) error {
	viper.SetConfigFile(configFilePath)
	return viper.ReadInConfig() // Find and read the config file
}
