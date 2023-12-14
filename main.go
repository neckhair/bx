package main

import (
	"github.com/neckhair/bx/cmd"
	"github.com/neckhair/bx/config"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(config.ConfigFileFullPath)

	cmd.Execute()
}
