package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/spf13/viper"
)

const (
	ConfigFileName = "config.yaml"
)

var (
	ConfigDirectory    = filepath.Join(xdg.ConfigHome, "bx")
	ConfigFileFullPath = filepath.Join(ConfigDirectory, ConfigFileName)
)

func WriteToFile() error {
	if err := ensureDirectoryExists(); err != nil {
		return fmt.Errorf("cannot create config directory: %w", err)
	}
	if err := viper.WriteConfigAs(ConfigFileFullPath); err != nil {
		return fmt.Errorf("cannot write config file: %v", err)
	}
	return nil
}

func ensureDirectoryExists() error {
	if err := os.MkdirAll(ConfigDirectory, 0o700); err != nil {
		return err
	}
	return nil
}
