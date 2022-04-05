package config

import (
	"fmt"

	"github.com/spf13/viper"
)

//Constants
const (
	defaultEnv = "qa"
)

var cfgReader *configReader

type (
	Configuration struct {
		DatabaseSettings
		NewRelicSettings
		JwtSettings
	}

	DatabaseSettings struct {
		DatabaseURI  string
		DatabaseName string
		Username     string
		Password     string
	}
	NewRelicSettings struct {
		LicenseKey string
		AppName    string
	}
	JwtSettings struct {
		SecretKey string
	}
	configReader struct {
		configFile string
		v          *viper.Viper
	}
)

func GetAllConfigValues(configFile string) (configuration *Configuration, err error) {
	newConfigReader(configFile)
	if err = cfgReader.v.ReadInConfig(); err != nil {
		fmt.Printf("Failed to read config file : %s", err)
		return nil, err
	}

	err = cfgReader.v.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Failed to unmarshal yaml file to configuration struct : %s", err)
		return nil, err
	}

	return configuration, err
}

func newConfigReader(configFile string) {
	v := viper.GetViper()
	v.SetConfigType("yaml")
	v.SetConfigFile(configFile)
	cfgReader = &configReader{
		configFile: configFile,
		v:          v,
	}
}
