package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

const configFilename = "config.yaml"

var (
	configDirectory    = filepath.Join(xdg.ConfigHome, "bx")
	configFileFullPath = filepath.Join(configDirectory, configFilename)
)

var ErrConfigNotFound = errors.New("config file not found")

func Init() error {
	viper.AddConfigPath(configDirectory)
	viper.SetConfigFile(configFilename)
	viper.SetConfigType("yaml")

	if err := os.MkdirAll(configDirectory, 0o700); err != nil {
		return fmt.Errorf("fatal error creating config directory: %w", err)
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return ErrConfigNotFound
		} else {
			return fmt.Errorf("fatal error config file: %w", err)
		}
	}
	return nil
}

func Set(key string, value interface{}) {
	viper.Set(key, value)
}

func Write() error {
	if err := viper.WriteConfigAs(configFileFullPath); err != nil {
		return fmt.Errorf("writing config failed %w", err)
	}
	return nil
}
