package main

import (
	"github.com/aarthikrao/payLaterService/cli"
	"github.com/aarthikrao/payLaterService/config"
	conn "github.com/aarthikrao/payLaterService/connections"
)

func main() {
	// Initialise configs
	config.InitConfigs()

	// Initialise the postgres db connection
	conn.InitDatabase()
	// Close db connection at shutdown
	defer conn.ShutDown()

	// Initialise services used for payments
	cli.INITServices()

	// Start CLI
	cli.RunCli()
}
