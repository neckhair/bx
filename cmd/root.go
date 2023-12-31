package cmd

import (
	"fmt"
	"os"

	"github.com/neckhair/bx/config"
	"github.com/neckhair/bx/internal/cli"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bx",
	Short: "Bexio command line",
	Long:  `Unofficial command line client for Bexio.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var pathError *os.PathError

		if err := viper.ReadInConfig(); err != nil {
			if !errors.As(err, &viper.ConfigFileNotFoundError{}) && !errors.As(err, &pathError) {
				return fmt.Errorf("error reading configuration, %w", err)
			}
		}

		// no further config checks for the setup command
		if cmd == setupCmd {
			return nil
		}

		if !config.CredentialsPresent() {
			cli.PrintError("Missing Bexio credentials. Please use the setup command first.")
			return nil
		}

		token, err := config.ReadTokenFromFile()
		if err != nil {
			return errors.Wrap(err, "Reading authentication token failed.")
		}
		if token.AccessToken == "" {
			cli.PrintError("Missing authentication token. Please use the setup command first.")
			return nil
		}

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bx.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
