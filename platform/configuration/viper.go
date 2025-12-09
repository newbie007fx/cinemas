package configuration

import (
	"github.com/spf13/viper"
)

type ConfigService struct {
	path       string
	configName string
	configType string
	config     *Config
}

func New(path, configName, configType string) *ConfigService {
	return &ConfigService{
		path:       path,
		configName: configName,
		configType: configType,
		config:     &Config{},
	}
}

func (cs *ConfigService) Setup() (err error) {
	viper.SetConfigName(cs.configName)
	viper.SetConfigType(cs.configType)
	viper.AddConfigPath(cs.path)

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(cs.config)

	return
}

func (cs *ConfigService) GetConfig() *Config {
	return cs.config
}
