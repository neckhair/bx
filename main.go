package main

import (
	"github.com/spf13/viper"
	"gitlab.com/neckhair/bx/cmd"
	"gitlab.com/neckhair/bx/config"
)

func main() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(config.ConfigFileFullPath)

	cmd.Execute()
}
