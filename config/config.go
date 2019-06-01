package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config is the global variable for configs
var Config *viper.Viper

// InitConfigs is used to initialise the configs
func InitConfigs() {
	Config = viper.New()
	Config.SetConfigFile("config.json")
	err := Config.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("Initialised configs")
}
