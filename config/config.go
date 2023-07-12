package config

import "github.com/spf13/viper"

// LoadConfig loads the conf.yaml from a given path passed to it
func LoadConfig(configPath string) error {
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)
	return viper.ReadInConfig() // Find and read the config file
}
