package config

import (
	"bytes"
	"fmt"
	"log"

	"github.com/gobuffalo/packr/v2"
	"github.com/spf13/viper"
)

//Configuration config struct
type Configuration struct {
	Nats     nats
}

type nats struct {
	URL string
}

// AppEnvironment represents global app environment
var AppEnvironment interface{}

//AppConfig app configuratiuon
var AppConfig Configuration

// InitConfiguration init config with viper
func InitConfiguration(configBox *packr.Box) Configuration {
	viper.SetEnvPrefix("go")
	viper.BindEnv("env")
	viper.SetConfigType("toml")
	AppEnvironment = viper.Get("env")
	if AppEnvironment == nil {
		AppEnvironment = "development"
	}
	config, err := configBox.Find(fmt.Sprintf("config.%v.toml", AppEnvironment))
	if err != nil {
		log.Fatal(err)
	}
	err = viper.ReadConfig(bytes.NewBuffer(config))
	if err != nil {
		log.Fatal(err)
	}
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}
	return AppConfig
}

// IsEnvironment checks app environment
func IsEnvironment(env string) bool {
	return AppEnvironment == env
}

// IsDevelopment checks app environment is development
func IsDevelopment() bool {
	return IsEnvironment("development")
}