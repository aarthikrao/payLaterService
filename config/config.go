package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Config is the global variable for configs
var Config *viper.Viper

// InitConfigs is used to initialise the configs
func InitConfigs() {
	Config = viper.New()
	home := os.Getenv("HOME")
	if home == "" {
		fmt.Println("Please set $HOME path")
	}
	// fmt.Println("Path home:", home)
	Config.SetConfigFile(home + "/configs/config.json")

	err := Config.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("Initialised configs")
}
